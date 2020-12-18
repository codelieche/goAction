package rpcdemo

import "github.com/pkg/errors"

type DemoService struct {
	X, Y int
}

type Args struct {
	X, Y int
}

func (DemoService) Div (args Args, result *float64) error {
	if args.Y == 0 {
		return errors.New("参数错误")
	}
	*result = float64(args.X) / float64(args.Y)
	return nil
}


