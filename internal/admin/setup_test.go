package admin_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/jinzhu/gorm"

	"github.com/qorpress/qorpress/internal/admin"
	. "github.com/qorpress/qorpress/internal/admin/tests/dummy"
)

var (
	server *httptest.Server
	db     *gorm.DB
	Admin  *admin.Admin
)

func init() {
	var mux http.Handler
	mux, Admin, db = NewTestHandler()
	server = httptest.NewServer(mux)
}

func NewTestHandler() (h http.Handler, adm *admin.Admin, d *gorm.DB) {
	adm = NewDummyAdmin()
	d = adm.DB
	h = adm.NewServeMux("/admin")
	return
}
