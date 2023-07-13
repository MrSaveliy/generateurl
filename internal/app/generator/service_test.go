package generator

import (
	"context"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"name/generator/pkg/api/generator"
	"os"
	"regexp"
	"testing"
)

var (
	rx = regexp.MustCompile(`^[A-Za-z0-9]+$`)
)

type testSuite struct {
	suite.Suite
	i *Implementation
}

func TestGenerator(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(testSuite))
}

func (s *testSuite) SetupSuite() {
	s.Require().NoError(os.Setenv("DB_TYPE", "in-memory"))
	s.i = NewImplementation()
}

func (s *testSuite) TestGetUrl() {
	s.Require().NoError(s.i.r.Set("url", "short"))

	ctx := context.TODO()

	// not found
	_, err := s.i.GetUrl(ctx, &generator.GetUrlRequest{Short: "short1"})
	s.Require().Equal(codes.NotFound, status.Code(err))

	// ok
	resp, err := s.i.GetUrl(ctx, &generator.GetUrlRequest{Short: "short"})
	s.Require().NoError(err)
	s.Require().Equal("url", resp.GetUrl())
}

func (s *testSuite) TestSetUrl() {
	resp, err := s.i.SetUrl(context.TODO(), &generator.SetUrlRequest{Url: "url"})
	s.Require().NoError(err)
	s.Require().Len(resp.GetShort(), 10)
	s.Require().True(rx.MatchString(resp.GetShort()))

	short, err := s.i.GetUrl(context.TODO(), &generator.GetUrlRequest{Short: resp.GetShort()})
	s.Require().NoError(err)
	s.Require().Equal("url", short.Url)
}
