package storage

import (
	p "public_service/genproto/public"
)



type PartyI interface{
	Create(*p.CreatePartyReq)(*p.Void, error)
	GetById(*p.ById)(*p.GetPartyRes, error)
	GetAll(*p.Filter)(*p.GetAllPartyRes, error)
	Update(*p.GetPartyRes)(*p.Void, error)
	Delete(*p.ById)(*p.Void, error)
}

type PublicI interface{
	Create(*p.CreatePublicReq)(*p.Void, error)
	GetById(*p.ById)(*p.GetPublicRes, error)
	GetAll(*p.Filter)(*p.GetAllPublicRes, error)
	Update(*p.GetPublicRes)(*p.Void, error)
	Delete(*p.ById)(*p.Void, error)
}