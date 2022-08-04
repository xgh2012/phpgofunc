package utilsjson

import "testing"

type St struct {
	Stkey1 string `json:"stkey_1"`
	Stkey2 int
	Stkey3 bool
	Stkey4 *St `json:"stkey_4,omitempty"`
}

func initData() (interface{}, map[string]interface{}) {
	a := true
	var data interface{}
	var data1 map[string]interface{}
	var i bool = false
	var i1 bool = true

	dataTmp := map[string]interface{}{
		"key2":  "2",
		"key4":  `{"barid":"44030610001028","tel":"1301112","open":1}`,
		"key8":  `{}`,
		"key9":  `[]`,
		"key3":  true,
		"key10": false,
		"key1":  1,
		"key5": map[string]interface{}{
			"k51":   "mst",
			"barid": "44030610001028",
			"k53":   64,
			"key521": map[string]interface{}{
				"key5211": "mst",
				"key5212": "44030610001028",
				"key5213": 64,
			},
		},
		"key11": map[string]int{
			"k111": 111,
			"k112": 112,
			"k113": 113,
		},
		"key16": map[string]bool{
			"k161": false,
			"k162": true,
			"k163": true,
		},
		"key13": map[string]string{
			"k131": "131",
			"k132": "132",
			"k133": "133",
		},
		"key12": []*bool{
			&i,
			&i1,
		},
		"key14": []St{
			St{
				Stkey1: "key141",
				Stkey2: 141,
				Stkey3: false,
			},
			St{
				Stkey1: "key142",
				Stkey2: 142,
				Stkey3: true,
			},
		},
		"key17": []*St{
			&St{
				Stkey1: "key171",
				Stkey2: 171,
				Stkey3: false,
			},
			&St{
				Stkey1: "key172",
				Stkey2: 172,
				Stkey3: true,
				Stkey4: &St{
					Stkey1: "key1721",
					Stkey2: 1721,
					Stkey3: true,
				},
			},
		},
		"key6": St{
			Stkey1: "k61",
			Stkey2: 62,
			Stkey3: false,
			Stkey4: &St{
				Stkey1: "k641",
				Stkey2: 641,
				Stkey3: false,
			},
		},
		"key7": &St{
			Stkey1: "k71",
			Stkey2: 72,
			Stkey3: true,
		},
		"key15": &a,
	}
	data = dataTmp
	data1 = dataTmp
	return data, data1
}
func TestHandleAnyToString(t *testing.T) {
	type args struct {
		data interface{}
	}
	data, _ := initData()
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test AnyToString",
			args: args{
				data: data,
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HandleAnyToString(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("HandleAnyToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HandleAnyToString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandleMapAnyToString(t *testing.T) {
	type args struct {
		data map[string]interface{}
	}
	_, data := initData()
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test MapAnyToString",
			args: args{
				data: data,
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HandleMapAnyToString(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("HandleMapAnyToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HandleMapAnyToString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
