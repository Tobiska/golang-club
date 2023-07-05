package concurency

/*
	Реавлизовать функцию pumpLoader, которая копит значения из канала in,
	отправляет обработанные сообщения batch'ами в канал out, если размер batch'а превысил L
	или время ожидания превысило T.
*/

const (
	L = 20
	T = 3000
)

type Output struct {
	processed []string
}

func pumpLoader(in chan string, out chan Output) {
}

func main() {
	in := make(chan string)

	out := make(chan Output)

	go pumpLoader(in, out)

	//...
}
