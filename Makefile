generate-protos:
	buf breaking --against ".git#branch=main"
	buf generate
