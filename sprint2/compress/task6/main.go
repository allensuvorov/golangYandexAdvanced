// LengthHandle возвращает размер распакованных данных.
func LengthHandle(w http.ResponseWriter, r *http.Request) {
	// создаём *gzip.Reader, который будет читать тело запроса
	// и распаковывать его
	gz, err := gzip.NewReader(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// не забывайте потом закрыть *gzip.Reader
	defer gz.Close()

	// при чтении вернётся распакованный слайс байт
	body, err := io.ReadAll(gz)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Length: %d", len(body))
}