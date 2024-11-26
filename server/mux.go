// Copyright (C) 2024 Pavel Sobolev
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package server

import (
	"net/http"

	"morbo/db"
)

type ServeMux struct {
	http.ServeMux

	feedHandler    feedHandler
	sessionHandler sessionHandler
}

func NewServeMux(db *db.DB) *ServeMux {
	mux := ServeMux{
		feedHandler:    feedHandler{db},
		sessionHandler: sessionHandler{db},
	}

	mux.Handle("/{$}", http.NotFoundHandler())
	mux.Handle("/feed/{$}", &mux.feedHandler)
	mux.Handle("/session/{$}", &mux.sessionHandler)

	return &mux
}
