package src

import (
	"api/main/src/bclient"
	privat "api/main/src/bclient/privat_bank/impl"
)

type Dependencies struct {
	ClientCore *bclient.Core
}

func createBClientCore() *bclient.Core {
	core := bclient.NewBClientCore()
	core.AddClient("privat_bank", privat.NewClient("Test token"))
	return core
}

func NewDependencies() (*Dependencies, func()) {
	core := createBClientCore()

	return &Dependencies{
			ClientCore: core,
		}, func() {

		}
}
