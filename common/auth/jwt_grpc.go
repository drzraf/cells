/*
 * Copyright (c) 2018. Abstrium SAS <team (at) pydio.com>
 * This file is part of Pydio Cells.
 *
 * Pydio Cells is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Pydio Cells is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Pydio Cells.  If not, see <http://www.gnu.org/licenses/>.
 *
 * The latest code can be found at <https://pydio.com>.
 */

package auth

import (
	"context"

	"github.com/mitchellh/mapstructure"
	"github.com/ory/fosite/token/jwt"
	"golang.org/x/oauth2"

	"github.com/pydio/cells/common/auth/claim"
	"github.com/pydio/cells/common/auth/hydra"
	defaults "github.com/pydio/cells/common/micro"
	"github.com/pydio/cells/common/proto/auth"
	json "github.com/pydio/cells/x/jsonx"
)

type grpcprovider struct {
	service string
}

type grpctoken struct {
	claims *jwt.IDTokenClaims
}

func RegisterGRPCProvider(service string) {
	p := new(grpcprovider)

	p.service = service

	addProvider(p)
}

func (p *grpcprovider) GetType() ProviderType {
	return ProviderTypeGrpc
}

func (c *grpcprovider) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	return hydra.Exchange(ctx, code)
}

func (c *grpcprovider) Verify(ctx context.Context, rawIDToken string) (IDToken, error) {

	cli := auth.NewAuthTokenVerifierClient(c.service, defaults.NewClient())

	resp, err := cli.Verify(ctx, &auth.VerifyTokenRequest{
		Token: rawIDToken,
	})
	if err != nil {
		return nil, err
	}

	token := new(grpctoken)

	if err := json.Unmarshal(resp.GetData(), &token.claims); err != nil {
		return nil, err
	}

	return token, nil
}

func (t *grpctoken) Claims(v interface{}) error {
	return mapstructure.Decode(t.claims.ToMap(), &v)
}

func (t *grpctoken) ScopedClaims(claims *claim.Claims) error {
	if ss := t.claims.Get("scopes"); ss != nil {
		m, _ := json.Marshal(ss)
		var parsed []string
		if e := json.Unmarshal(m, &parsed); e == nil {
			claims.ProvidesScopes = true
			claims.Scopes = append(claims.Scopes, parsed...)
		} else {
			return e
		}
	}
	return nil
}
