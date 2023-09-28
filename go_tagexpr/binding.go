/*
 * Copyright 2023 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package go_tagexpr

import (
	"encoding/json"
	"reflect"

	"github.com/bytedance/go-tagexpr/v2/binding"
	"github.com/bytedance/go-tagexpr/v2/binding/gjson"
	"github.com/bytedance/go-tagexpr/v2/validator"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/route/param"
)

func init() {
	binding.ResetJSONUnmarshaler(sonic.Unmarshal)
}

var defaultBinder = NewBinder()

type Binder struct {
	*binding.Binding
}

func NewBinder() *Binder {
	return &Binder{
		Binding: binding.Default(),
	}
}

func (b *Binder) Name() string {
	return "go-tagexpr"
}

func (b *Binder) Bind(request *protocol.Request, i interface{}, params param.Params) error {
	return b.Binding.IBind(i, wrapRequest(request), params)
}

func (b *Binder) BindAndValidate(request *protocol.Request, i interface{}, params param.Params) error {
	return b.Binding.IBindAndValidate(i, wrapRequest(request), params)
}

func (b *Binder) BindQuery(request *protocol.Request, i interface{}) error {
	panic("unsupport BindQuery")
}

func (b *Binder) BindHeader(request *protocol.Request, i interface{}) error {
	panic("unsupport BindHeader")
}

func (b *Binder) BindPath(request *protocol.Request, i interface{}, params param.Params) error {
	panic("unsupport BindPath")
}

func (b *Binder) BindForm(request *protocol.Request, i interface{}) error {
	panic("unsupport BindForm")
}

func (b *Binder) BindJSON(request *protocol.Request, i interface{}) error {
	panic("unsupport BindJSON")
}

func (b *Binder) BindProtobuf(request *protocol.Request, i interface{}) error {
	panic("unsupport BindProtobuf")
}

// BindAndValidate binds data from *protocol.Request to obj and validates them if needed.
// NOTE:
//
//	obj should be a pointer.
func BindAndValidate(req *protocol.Request, obj interface{}, pathParams param.Params) error {
	return defaultBinder.BindAndValidate(req, obj, pathParams)
}

// Bind binds data from *protocol.Request to obj.
// NOTE:
//
//	obj should be a pointer.
func Bind(req *protocol.Request, obj interface{}, pathParams param.Params) error {
	return defaultBinder.Bind(req, obj, pathParams)
}

// SetLooseZeroMode if set to true,
// the empty string request parameter is bound to the zero value of parameter.
// NOTE:
//
//	The default is false.
//	Suitable for these parameter types: query/header/cookie/form .
func SetLooseZeroMode(enable bool) {
	defaultBinder.SetLooseZeroMode(enable)
}

// SetErrorFactory customizes the factory of validation error.
// NOTE:
//
//	If errFactory==nil, the default is used.
//	SetErrorFactory will remain in effect once it has been called.
func SetErrorFactory(bindErrFactory, validatingErrFactory func(failField, msg string) error) {
	defaultBinder.SetErrorFactory(bindErrFactory, validatingErrFactory)
}

// MustRegTypeUnmarshal registers unmarshal function of type.
// NOTE:
//
//	It will panic if exist error.
//	MustRegTypeUnmarshal will remain in effect once it has been called.
func MustRegTypeUnmarshal(t reflect.Type, fn func(v string, emptyAsZero bool) (reflect.Value, error)) {
	binding.MustRegTypeUnmarshal(t, fn)
}

// MustRegValidateFunc registers validator function expression.
// NOTE:
//
//	If force=true, allow to cover the existed same funcName.
//	MustRegValidateFunc will remain in effect once it has been called.
func MustRegValidateFunc(funcName string, fn func(args ...interface{}) error, force ...bool) {
	validator.RegFunc(funcName, fn, force...)
}

// UseStdJSONUnmarshaler uses encoding/json as json library
// NOTE:
//
//	The current version uses encoding/json by default.
//	UseStdJSONUnmarshaler will remain in effect once it has been called.
func UseStdJSONUnmarshaler() {
	binding.ResetJSONUnmarshaler(json.Unmarshal)
}

// UseGJSONUnmarshaler uses github.com/bytedance/go-tagexpr/v2/binding/gjson as json library
// NOTE:
//
//	UseGJSONUnmarshaler will remain in effect once it has been called.
func UseGJSONUnmarshaler() {
	gjson.UseJSONUnmarshaler()
}

// UseThirdPartyJSONUnmarshaler uses third-party json library for binding
// NOTE:
//
//	UseThirdPartyJSONUnmarshaler will remain in effect once it has been called.
func UseThirdPartyJSONUnmarshaler(unmarshaler func(data []byte, v interface{}) error) {
	binding.ResetJSONUnmarshaler(unmarshaler)
}

// Validate validates obj with "vd" tag
// NOTE:
//
//	obj should be a pointer.
//	Validate should be called after Bind.
func Validate(obj interface{}) error {
	return defaultValidator.Validate(obj)
}

var defaultValidator = NewValidator()

type Validator struct {
	*binding.Binding
}

func NewValidator() *Validator {
	return &Validator{
		Binding: binding.Default(),
	}
}

func (v *Validator) ValidateStruct(i interface{}) error {
	return v.Validate(i)
}

func (v *Validator) Engine() interface{} {
	return v.Binding
}

func (v *Validator) ValidateTag() string {
	return "vd"
}
