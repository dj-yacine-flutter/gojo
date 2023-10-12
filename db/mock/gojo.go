// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dj-yacine-flutter/gojo/db/database (interfaces: Gojo)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	db "github.com/dj-yacine-flutter/gojo/db/database"
	gomock "github.com/golang/mock/gomock"
	pgtype "github.com/jackc/pgx/v5/pgtype"
)

// MockGojo is a mock of Gojo interface.
type MockGojo struct {
	ctrl     *gomock.Controller
	recorder *MockGojoMockRecorder
}

// MockGojoMockRecorder is the mock recorder for MockGojo.
type MockGojoMockRecorder struct {
	mock *MockGojo
}

// NewMockGojo creates a new mock instance.
func NewMockGojo(ctrl *gomock.Controller) *MockGojo {
	mock := &MockGojo{ctrl: ctrl}
	mock.recorder = &MockGojoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGojo) EXPECT() *MockGojoMockRecorder {
	return m.recorder
}

// CreateAnimeGenre mocks base method.
func (m *MockGojo) CreateAnimeGenre(arg0 context.Context, arg1 db.CreateAnimeGenreParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAnimeGenre", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAnimeGenre indicates an expected call of CreateAnimeGenre.
func (mr *MockGojoMockRecorder) CreateAnimeGenre(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAnimeGenre", reflect.TypeOf((*MockGojo)(nil).CreateAnimeGenre), arg0, arg1)
}

// CreateAnimeMeta mocks base method.
func (m *MockGojo) CreateAnimeMeta(arg0 context.Context, arg1 db.CreateAnimeMetaParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAnimeMeta", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAnimeMeta indicates an expected call of CreateAnimeMeta.
func (mr *MockGojoMockRecorder) CreateAnimeMeta(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAnimeMeta", reflect.TypeOf((*MockGojo)(nil).CreateAnimeMeta), arg0, arg1)
}

// CreateAnimeMovie mocks base method.
func (m *MockGojo) CreateAnimeMovie(arg0 context.Context, arg1 db.CreateAnimeMovieParams) (db.AnimeMovie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAnimeMovie", arg0, arg1)
	ret0, _ := ret[0].(db.AnimeMovie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAnimeMovie indicates an expected call of CreateAnimeMovie.
func (mr *MockGojoMockRecorder) CreateAnimeMovie(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAnimeMovie", reflect.TypeOf((*MockGojo)(nil).CreateAnimeMovie), arg0, arg1)
}

// CreateAnimeMovieTx mocks base method.
func (m *MockGojo) CreateAnimeMovieTx(arg0 context.Context, arg1 db.CreateAnimeMovieTxParams) (db.CreateAnimeMovieTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAnimeMovieTx", arg0, arg1)
	ret0, _ := ret[0].(db.CreateAnimeMovieTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAnimeMovieTx indicates an expected call of CreateAnimeMovieTx.
func (mr *MockGojoMockRecorder) CreateAnimeMovieTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAnimeMovieTx", reflect.TypeOf((*MockGojo)(nil).CreateAnimeMovieTx), arg0, arg1)
}

// CreateAnimeStudio mocks base method.
func (m *MockGojo) CreateAnimeStudio(arg0 context.Context, arg1 db.CreateAnimeStudioParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAnimeStudio", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAnimeStudio indicates an expected call of CreateAnimeStudio.
func (mr *MockGojoMockRecorder) CreateAnimeStudio(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAnimeStudio", reflect.TypeOf((*MockGojo)(nil).CreateAnimeStudio), arg0, arg1)
}

// CreateGenre mocks base method.
func (m *MockGojo) CreateGenre(arg0 context.Context, arg1 string) (db.Genre, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGenre", arg0, arg1)
	ret0, _ := ret[0].(db.Genre)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGenre indicates an expected call of CreateGenre.
func (mr *MockGojoMockRecorder) CreateGenre(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGenre", reflect.TypeOf((*MockGojo)(nil).CreateGenre), arg0, arg1)
}

// CreateLanguage mocks base method.
func (m *MockGojo) CreateLanguage(arg0 context.Context, arg1 db.CreateLanguageParams) (db.Language, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLanguage", arg0, arg1)
	ret0, _ := ret[0].(db.Language)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLanguage indicates an expected call of CreateLanguage.
func (mr *MockGojoMockRecorder) CreateLanguage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLanguage", reflect.TypeOf((*MockGojo)(nil).CreateLanguage), arg0, arg1)
}

// CreateMeta mocks base method.
func (m *MockGojo) CreateMeta(arg0 context.Context, arg1 db.CreateMetaParams) (db.Meta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMeta", arg0, arg1)
	ret0, _ := ret[0].(db.Meta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMeta indicates an expected call of CreateMeta.
func (mr *MockGojoMockRecorder) CreateMeta(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMeta", reflect.TypeOf((*MockGojo)(nil).CreateMeta), arg0, arg1)
}

// CreateStudio mocks base method.
func (m *MockGojo) CreateStudio(arg0 context.Context, arg1 string) (db.Studio, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStudio", arg0, arg1)
	ret0, _ := ret[0].(db.Studio)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStudio indicates an expected call of CreateStudio.
func (mr *MockGojoMockRecorder) CreateStudio(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStudio", reflect.TypeOf((*MockGojo)(nil).CreateStudio), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockGojo) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockGojoMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockGojo)(nil).CreateUser), arg0, arg1)
}

// DeleteAnimeGenre mocks base method.
func (m *MockGojo) DeleteAnimeGenre(arg0 context.Context, arg1 db.DeleteAnimeGenreParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAnimeGenre", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAnimeGenre indicates an expected call of DeleteAnimeGenre.
func (mr *MockGojoMockRecorder) DeleteAnimeGenre(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAnimeGenre", reflect.TypeOf((*MockGojo)(nil).DeleteAnimeGenre), arg0, arg1)
}

// DeleteAnimeMeta mocks base method.
func (m *MockGojo) DeleteAnimeMeta(arg0 context.Context, arg1 db.DeleteAnimeMetaParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAnimeMeta", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAnimeMeta indicates an expected call of DeleteAnimeMeta.
func (mr *MockGojoMockRecorder) DeleteAnimeMeta(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAnimeMeta", reflect.TypeOf((*MockGojo)(nil).DeleteAnimeMeta), arg0, arg1)
}

// DeleteAnimeMovie mocks base method.
func (m *MockGojo) DeleteAnimeMovie(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAnimeMovie", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAnimeMovie indicates an expected call of DeleteAnimeMovie.
func (mr *MockGojoMockRecorder) DeleteAnimeMovie(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAnimeMovie", reflect.TypeOf((*MockGojo)(nil).DeleteAnimeMovie), arg0, arg1)
}

// DeleteAnimeStudio mocks base method.
func (m *MockGojo) DeleteAnimeStudio(arg0 context.Context, arg1 db.DeleteAnimeStudioParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAnimeStudio", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAnimeStudio indicates an expected call of DeleteAnimeStudio.
func (mr *MockGojoMockRecorder) DeleteAnimeStudio(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAnimeStudio", reflect.TypeOf((*MockGojo)(nil).DeleteAnimeStudio), arg0, arg1)
}

// DeleteGenre mocks base method.
func (m *MockGojo) DeleteGenre(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGenre", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGenre indicates an expected call of DeleteGenre.
func (mr *MockGojoMockRecorder) DeleteGenre(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGenre", reflect.TypeOf((*MockGojo)(nil).DeleteGenre), arg0, arg1)
}

// DeleteLanguage mocks base method.
func (m *MockGojo) DeleteLanguage(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLanguage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLanguage indicates an expected call of DeleteLanguage.
func (mr *MockGojoMockRecorder) DeleteLanguage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLanguage", reflect.TypeOf((*MockGojo)(nil).DeleteLanguage), arg0, arg1)
}

// DeleteMeta mocks base method.
func (m *MockGojo) DeleteMeta(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMeta", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMeta indicates an expected call of DeleteMeta.
func (mr *MockGojoMockRecorder) DeleteMeta(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMeta", reflect.TypeOf((*MockGojo)(nil).DeleteMeta), arg0, arg1)
}

// DeleteStudio mocks base method.
func (m *MockGojo) DeleteStudio(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStudio", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStudio indicates an expected call of DeleteStudio.
func (mr *MockGojoMockRecorder) DeleteStudio(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStudio", reflect.TypeOf((*MockGojo)(nil).DeleteStudio), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockGojo) DeleteUser(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockGojoMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockGojo)(nil).DeleteUser), arg0, arg1)
}

// GetAnimeMovie mocks base method.
func (m *MockGojo) GetAnimeMovie(arg0 context.Context, arg1 int64) (db.AnimeMovie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAnimeMovie", arg0, arg1)
	ret0, _ := ret[0].(db.AnimeMovie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAnimeMovie indicates an expected call of GetAnimeMovie.
func (mr *MockGojoMockRecorder) GetAnimeMovie(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAnimeMovie", reflect.TypeOf((*MockGojo)(nil).GetAnimeMovie), arg0, arg1)
}

// GetMetaIDByAnimeAndLanguage mocks base method.
func (m *MockGojo) GetMetaIDByAnimeAndLanguage(arg0 context.Context, arg1 db.GetMetaIDByAnimeAndLanguageParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetaIDByAnimeAndLanguage", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetaIDByAnimeAndLanguage indicates an expected call of GetMetaIDByAnimeAndLanguage.
func (mr *MockGojoMockRecorder) GetMetaIDByAnimeAndLanguage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetaIDByAnimeAndLanguage", reflect.TypeOf((*MockGojo)(nil).GetMetaIDByAnimeAndLanguage), arg0, arg1)
}

// GetUserByID mocks base method.
func (m *MockGojo) GetUserByID(arg0 context.Context, arg1 int64) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockGojoMockRecorder) GetUserByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockGojo)(nil).GetUserByID), arg0, arg1)
}

// GetUserByUsername mocks base method.
func (m *MockGojo) GetUserByUsername(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername.
func (mr *MockGojoMockRecorder) GetUserByUsername(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockGojo)(nil).GetUserByUsername), arg0, arg1)
}

// ListAnimeGenres mocks base method.
func (m *MockGojo) ListAnimeGenres(arg0 context.Context, arg1 db.ListAnimeGenresParams) ([]pgtype.Int4, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAnimeGenres", arg0, arg1)
	ret0, _ := ret[0].([]pgtype.Int4)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAnimeGenres indicates an expected call of ListAnimeGenres.
func (mr *MockGojoMockRecorder) ListAnimeGenres(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAnimeGenres", reflect.TypeOf((*MockGojo)(nil).ListAnimeGenres), arg0, arg1)
}

// ListAnimeMetas mocks base method.
func (m *MockGojo) ListAnimeMetas(arg0 context.Context, arg1 db.ListAnimeMetasParams) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAnimeMetas", arg0, arg1)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAnimeMetas indicates an expected call of ListAnimeMetas.
func (mr *MockGojoMockRecorder) ListAnimeMetas(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAnimeMetas", reflect.TypeOf((*MockGojo)(nil).ListAnimeMetas), arg0, arg1)
}

// ListAnimeStudios mocks base method.
func (m *MockGojo) ListAnimeStudios(arg0 context.Context, arg1 db.ListAnimeStudiosParams) ([]pgtype.Int4, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAnimeStudios", arg0, arg1)
	ret0, _ := ret[0].([]pgtype.Int4)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAnimeStudios indicates an expected call of ListAnimeStudios.
func (mr *MockGojoMockRecorder) ListAnimeStudios(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAnimeStudios", reflect.TypeOf((*MockGojo)(nil).ListAnimeStudios), arg0, arg1)
}

// ListGenres mocks base method.
func (m *MockGojo) ListGenres(arg0 context.Context, arg1 db.ListGenresParams) ([]db.Genre, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGenres", arg0, arg1)
	ret0, _ := ret[0].([]db.Genre)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGenres indicates an expected call of ListGenres.
func (mr *MockGojoMockRecorder) ListGenres(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGenres", reflect.TypeOf((*MockGojo)(nil).ListGenres), arg0, arg1)
}

// ListLanguages mocks base method.
func (m *MockGojo) ListLanguages(arg0 context.Context, arg1 db.ListLanguagesParams) ([]db.Language, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLanguages", arg0, arg1)
	ret0, _ := ret[0].([]db.Language)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLanguages indicates an expected call of ListLanguages.
func (mr *MockGojoMockRecorder) ListLanguages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLanguages", reflect.TypeOf((*MockGojo)(nil).ListLanguages), arg0, arg1)
}

// ListStudios mocks base method.
func (m *MockGojo) ListStudios(arg0 context.Context, arg1 db.ListStudiosParams) ([]db.Studio, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStudios", arg0, arg1)
	ret0, _ := ret[0].([]db.Studio)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStudios indicates an expected call of ListStudios.
func (mr *MockGojoMockRecorder) ListStudios(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStudios", reflect.TypeOf((*MockGojo)(nil).ListStudios), arg0, arg1)
}

// ListUsers mocks base method.
func (m *MockGojo) ListUsers(arg0 context.Context, arg1 db.ListUsersParams) ([]db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", arg0, arg1)
	ret0, _ := ret[0].([]db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockGojoMockRecorder) ListUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockGojo)(nil).ListUsers), arg0, arg1)
}

// UpdateAnimeGenre mocks base method.
func (m *MockGojo) UpdateAnimeGenre(arg0 context.Context, arg1 db.UpdateAnimeGenreParams) (db.AnimeGenre, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAnimeGenre", arg0, arg1)
	ret0, _ := ret[0].(db.AnimeGenre)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAnimeGenre indicates an expected call of UpdateAnimeGenre.
func (mr *MockGojoMockRecorder) UpdateAnimeGenre(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAnimeGenre", reflect.TypeOf((*MockGojo)(nil).UpdateAnimeGenre), arg0, arg1)
}

// UpdateAnimeMeta mocks base method.
func (m *MockGojo) UpdateAnimeMeta(arg0 context.Context, arg1 db.UpdateAnimeMetaParams) (db.AnimeMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAnimeMeta", arg0, arg1)
	ret0, _ := ret[0].(db.AnimeMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAnimeMeta indicates an expected call of UpdateAnimeMeta.
func (mr *MockGojoMockRecorder) UpdateAnimeMeta(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAnimeMeta", reflect.TypeOf((*MockGojo)(nil).UpdateAnimeMeta), arg0, arg1)
}

// UpdateAnimeStudio mocks base method.
func (m *MockGojo) UpdateAnimeStudio(arg0 context.Context, arg1 db.UpdateAnimeStudioParams) (db.AnimeStudio, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAnimeStudio", arg0, arg1)
	ret0, _ := ret[0].(db.AnimeStudio)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAnimeStudio indicates an expected call of UpdateAnimeStudio.
func (mr *MockGojoMockRecorder) UpdateAnimeStudio(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAnimeStudio", reflect.TypeOf((*MockGojo)(nil).UpdateAnimeStudio), arg0, arg1)
}

// UpdateGenre mocks base method.
func (m *MockGojo) UpdateGenre(arg0 context.Context, arg1 db.UpdateGenreParams) (db.Genre, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGenre", arg0, arg1)
	ret0, _ := ret[0].(db.Genre)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateGenre indicates an expected call of UpdateGenre.
func (mr *MockGojoMockRecorder) UpdateGenre(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGenre", reflect.TypeOf((*MockGojo)(nil).UpdateGenre), arg0, arg1)
}

// UpdateLanguage mocks base method.
func (m *MockGojo) UpdateLanguage(arg0 context.Context, arg1 db.UpdateLanguageParams) (db.Language, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLanguage", arg0, arg1)
	ret0, _ := ret[0].(db.Language)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateLanguage indicates an expected call of UpdateLanguage.
func (mr *MockGojoMockRecorder) UpdateLanguage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLanguage", reflect.TypeOf((*MockGojo)(nil).UpdateLanguage), arg0, arg1)
}

// UpdateMeta mocks base method.
func (m *MockGojo) UpdateMeta(arg0 context.Context, arg1 db.UpdateMetaParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMeta", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMeta indicates an expected call of UpdateMeta.
func (mr *MockGojoMockRecorder) UpdateMeta(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMeta", reflect.TypeOf((*MockGojo)(nil).UpdateMeta), arg0, arg1)
}

// UpdateStudio mocks base method.
func (m *MockGojo) UpdateStudio(arg0 context.Context, arg1 db.UpdateStudioParams) (db.Studio, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStudio", arg0, arg1)
	ret0, _ := ret[0].(db.Studio)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStudio indicates an expected call of UpdateStudio.
func (mr *MockGojoMockRecorder) UpdateStudio(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStudio", reflect.TypeOf((*MockGojo)(nil).UpdateStudio), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockGojo) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockGojoMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockGojo)(nil).UpdateUser), arg0, arg1)
}
