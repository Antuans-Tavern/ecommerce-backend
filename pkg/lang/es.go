package lang

func setUpEs() {
	es, _ := translator.GetTranslator("es")

	es.Add("internal error", "Ha ocurrido un error en el servidor.", true)
	es.Add("login error", "Usuario o contraseña incorrectos.", true)
}
