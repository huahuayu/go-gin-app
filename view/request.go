package view

type (
	RegisterReq struct {
		Email    string
		Password string
		Username string
	}

	LoginReq struct {
		Email    string
		Password string
	}

	UpdatePasswordReq struct {
		OldPassword string
		NewPassword string
	}

	InfoReq struct{}

	UpdateUsernameReq struct {
		NewUsername string
	}

	AddParamReq struct {
		Type string
		Code string
		Desc string
	}

	GetParamByTypeReq struct {
		Type string
	}

	UpdateParamReq struct {
		Type string
		Code string
		Desc string
	}

	DeleteParamReq struct {
		Id string
	}
)
