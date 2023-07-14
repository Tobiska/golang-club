# Обработка ошибок

## С какими проблемами столкнулись?

1. Отсутствие классификации ошибок и их градации.
2. Множественная необоснованая реакция на ошибки(в основном логи).
3. Детализация ошибок.
4. Отсутствие трассировки стека в нужных нам ошибках.

## Реакция на ошибку
При возникновении ошибки должен быть описан код, который **явно** её обработает.

Возможные реакции:
- Логирование.
- Проброс выше по стеку вызова.
- Коллекционировать ошибку.
- Регистрация бага.
- Отправка уведомления (email, kafka и др.).

Следуя **бизнес процессу**, ошибка может быть обработана столько раз сколько это нужно.
Важно, что для реакции на ошибку нужна причина!

## Типы ошибок

В ходе работы часто сталкивались с тем, когда реакция на разные ошибочные кейсы реакция тоже должна быть разной.

### Было

```golang
func callingFunc(ctx context.Context) error {
    obj, isContinue, err := getObj(ctx context.Context)
    if err.String() == "user_error" {
        //...
    }

    if isContinue {
        //...
    }
}

func getObj(ctx context.Context, id string) (any, bool, error) {
   return fmt.Errorf("error")
}  
```

В пакете **makakierrors**(название рабочее) выделили 3 базовых и 3 расширенных типа ошибки(для пользовательских).
 - **UserError** (ошибки пользователя)
   - **UserCriticalError** (критические ошибки пользователя, чаще всего стоит прервать обработку чего либо)
   - **UserSoftError** (некритические, чаще всего стоит залогировать, но обработку продолжить)
 - **SystemError** (ошибки инфраструктурные)
 - **InternalError** (внутренние ошибки приложения, чаще всего баги)

```golang
makakierrors.WrapInternalError(errors.New("данный ID не содержится в собранных offerID"))
```

Так же можно создавать свои кастомные типы ошибок для конкретного пакета, унаследовавшись от одного из базовых.

Пример: 
 При обработке информации о товарах на складах, ошибочные кейсы требует особенной обработки:
 - при отсутствии id склада следует пропустить обработку скалада и не добавлять в итоговое сообщение (StoreSoftError).
 - при отсутсвии НДС, БЕИП и тд следует добавить в итоговое сообщение код ошибки (StoreCriticalError).

**StoreSoftError**, **StoreCriticalError** наследуются от **UserSoftError**.

```golang
for _, store := range stores {
    converter := NewStoreConverter(&store)
    storeDto, err := converter.Convert(c.offerDto.SupplierBasicUnit, offerErr)
    
    storesErr = makakierrors.Append(storesErr, err)
    if IsStoreCriticError(err) {
        continue
    }
    c.offerDto.Stores = append(c.offerDto.Stores, storeDto)
}
```

## Наследование

Типы ошибок могут "наследоваться" (вкладываться) друг в друга.

```golang
func WrapUserCriticError(err error) error {
	if err == nil {
		return nil
	}
	return &UserCriticError{
		error: WrapUserError(err),
	}
}
```

С помощью методов **errors.As** и **errors.Is** мы можем обходить всю вложенность в поисках конкретного типа или объекта ошибки.
В пакете **makakierrors** (название рабочее) поиск ошибки с конкретным типов происходит с помощью функций **IsUserCriticError**, **IsSystemError**

## Коллекционирование

На ошибку можно отреагировать, скопив её с остальными. Далее коллекцию ошибок можно вернуть под одним общим интерфейсом **error**.
Сделать это можно с помощью вспомогательных функций **Append** и **Combine** пакета **makakierrors** (название рабочее).

Пример:
```golang
func (c *OfferConverter) Convert(groupMap group.Map, countOffers int) (*dto.OfferDto, error) {
	//...

	c.offerDto = &dto.OfferDto{}

	convertError := makakierrors.Combine(
		c.FillSupplierProductId(),
		c.FillSupplierId(),
		//c.FillRatioBasicUnit(),
		c.FillBasicSupplierUnit(),
		c.FillVat(),
	)
	
	storesErr := c.FillStores(convertError)

	convertError = makakierrors.Append( //выглядит странненько)
		convertError,
		storesErr,
	)

	c.FillCountOffer(countOffers)
	c.FillRequestId()
	c.FillRequestTime()
	c.FillIsFull()

	return c.offerDto, convertError
}
```

Важно! Методы **IsUserCriticError**, **IsSystemError** и тд. способны находить ошибки определённого типа даже в коллекции ошибок.
Это огромный плюс данного подхода, но влечёт за собой проблемы производительности. Выше перечисленные методы будут обходить дерево ошибок в ширину пока не найдут нужный тип.

В данном случае стоит распологать проверки на ошибки по убыванию вложенности.

```golang
func (p *Processor) sendWrap(ctx context.Context, offerDto *dto.OfferDto) error {
    err := p.send(ctx, offerDto)
    switch {
        case err == nil:
            //...

		case makakierrors.IsUserCriticalError(err):
            //...
			
        case makakierrors.IsSystemError(err), makakierrors.IsInternalError(err):
            //...
}
```

## Обработка

```golang
func (p *Processor) sendWrap(ctx context.Context, offerDto *dto.OfferDto) error {
    err := p.send(ctx, offerDto)
    switch {
        case err == nil:
            log.L().Info(log.NewFromCtx(ctx).AddDetailAny(loggercontext.ResultKey, offerDto).
            AddMsg(loggercontext.KtpSendCode, fmt.Sprintf(loggercontext.KtpSendSuccessMsg, offerDto.Name, offerDto.Id, p.meta.SupplierId, p.meta.FilePath)))
        
        case makakierrors.IsSystemError(err), makakierrors.IsInternalError(err):
            log.L().Error(log.NewFromCtx(ctx).AddDetailAny(loggercontext.ResultKey, offerDto).AddBacktrace(err).
            AddMsg(loggercontext.KtpSendCode, fmt.Sprintf(loggercontext.KtpSendFailedMsg, offerDto.Name, offerDto.Id, p.meta.SupplierId, p.meta.FilePath)))
            return err
        
        case makakierrors.IsUserError(err):
            log.L().Warn(log.NewFromCtx(ctx).AddDetailAny(loggercontext.ResultKey, offerDto).AddBacktrace(err).
            AddMsg(loggercontext.KtpSendCode, fmt.Sprintf(loggercontext.KtpSendWarningMsg, offerDto.Name, offerDto.Id, p.meta.SupplierId, p.meta.FilePath)))
    }
    return nil
}
```
    
    


