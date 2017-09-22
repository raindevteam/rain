package hail

import "testing"

const (
	// Log.
	l = iota
	// Log with Error return.
	le
	// Log with format string.
	lf
	// Log with format string and Error return.
	lfe
)

type LogTestData struct {
	msg   string
	out   string
	vargs []interface{}
	t     int
}

func testLogFunction(t *testing.T, name string,
	ltd *LogTestData, v interface{}) {
	type args struct {
		facility int
		msg      string
		v        []interface{}
	}
	tests := []struct {
		name    string
		args    args
		output  string
		wantErr bool
	}{
		{
			name: name,
			args: args{
				facility: Fhail,
				msg:      ltd.msg,
				v:        ltd.vargs,
			},
			output:  ltd.out,
			wantErr: true,
		},
	}
	Defaults()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				switch ltd.t {
				case l:
					f := v.(func(int, string))
					f(tt.args.facility, tt.args.msg)
				case le:
					f := v.(func(int, string) error)
					if err := f(tt.args.facility, tt.args.msg); (err != nil) !=
						tt.wantErr {
						t.Errorf("Emerg() error = %v, wantErr = %v",
							err, tt.wantErr)
					}
				case lf:
					f := v.(func(int, string, ...interface{}))
					f(tt.args.facility, tt.args.msg, tt.args.v...)
				case lfe:
					f := v.(func(int, string, ...interface{}) error)
					err := f(tt.args.facility, tt.args.msg, tt.args.v...)
					if (err != nil) != tt.wantErr {
						t.Errorf("Emerg() error = %v, wantErr = %v",
							err, tt.wantErr)
					}
				}
			})
			if output != tt.output {
				t.Errorf("output = %s, want = %s", output, tt.output)
			}
		})
	}
}

func TestFatal(t *testing.T) {
	testLogFunction(t, "fatal log", &LogTestData{
		msg:   "unit 1",
		out:   "[HAIL] (1): unit 1\n",
		vargs: nil,
		t:     le,
	}, Fatal)
}

func TestFatalf(t *testing.T) {
	testLogFunction(t, "fatalf log", &LogTestData{
		msg:   "%s %d",
		out:   "[HAIL] (1): unit 2",
		vargs: []interface{}{"unit", 2},
		t:     lfe,
	}, Fatalf)
}

func TestCrit(t *testing.T) {
	testLogFunction(t, "crit log", &LogTestData{
		msg:   "unit 1",
		out:   "[HAIL] (2): unit 1\n",
		vargs: nil,
		t:     le,
	}, Crit)
}

func TestCritf(t *testing.T) {
	testLogFunction(t, "critf log", &LogTestData{
		msg:   "%s %d",
		out:   "[HAIL] (2): unit 2",
		vargs: []interface{}{"unit", 2},
		t:     lfe,
	}, Critf)
}

func TestErr(t *testing.T) {
	testLogFunction(t, "err log", &LogTestData{
		msg:   "unit 1",
		out:   "[HAIL] (3): unit 1\n",
		vargs: nil,
		t:     le,
	}, Err)
}

func TestErrf(t *testing.T) {
	testLogFunction(t, "errf log", &LogTestData{
		msg:   "%s %d",
		out:   "[HAIL] (3): unit 2",
		vargs: []interface{}{"unit", 2},
		t:     lfe,
	}, Errf)
}

func TestWarn(t *testing.T) {
	testLogFunction(t, "warn log", &LogTestData{
		msg:   "unit 1",
		out:   "[HAIL] (4): unit 1\n",
		vargs: nil,
		t:     l,
	}, Warn)
}

func TestWarnf(t *testing.T) {
	testLogFunction(t, "warnf log", &LogTestData{
		msg:   "%s %d",
		out:   "[HAIL] (4): unit 2",
		vargs: []interface{}{"unit", 2},
		t:     lf,
	}, Warnf)
}

func TestInfo(t *testing.T) {
	testLogFunction(t, "info log", &LogTestData{
		msg:   "unit 1",
		out:   "[HAIL] (5): unit 1\n",
		vargs: nil,
		t:     l,
	}, Info)
}

func TestInfof(t *testing.T) {
	testLogFunction(t, "infof log", &LogTestData{
		msg:   "%s %d",
		out:   "[HAIL] (5): unit 2",
		vargs: []interface{}{"unit", 2},
		t:     lf,
	}, Infof)
}

func TestDebug(t *testing.T) {
	testLogFunction(t, "debug log", &LogTestData{
		msg:   "unit 1",
		out:   "[HAIL] (6): unit 1\n",
		vargs: nil,
		t:     l,
	}, Debug)
}

func TestDebugf(t *testing.T) {
	testLogFunction(t, "debugf log", &LogTestData{
		msg:   "%s %d",
		out:   "[HAIL] (6): unit 2",
		vargs: []interface{}{"unit", 2},
		t:     lf,
	}, Debugf)
}
