// Copyright 2022 The Casdoor Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cred

import "github.com/alexedwards/argon2id"

type Argon2idCredManager struct{}

func NewArgon2idCredManager() *Argon2idCredManager {
	cm := &Argon2idCredManager{}
	return cm
}

func (cm *Argon2idCredManager) GetHashedPassword(password string, userSalt string, organizationSalt string) string {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return ""
	}
	return hash
}

func (cm *Argon2idCredManager) IsPasswordCorrect(plainPwd string, hashedPwd string, userSalt string, organizationSalt string) bool {
	match, _ := argon2id.ComparePasswordAndHash(plainPwd, hashedPwd)
	return match
}
