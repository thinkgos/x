package extstr

import (
	"testing"
)

func TestRecombine(t *testing.T) {
	type args struct {
		str string
		sep byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"空字符",
			args{
				str: "",
				sep: '_',
			},
			"",
		},
		{
			"大小写",
			args{
				str: "HelloWorld",
				sep: '_',
			},
			"hello_world",
		},
		{
			"大小写并带分隔符",
			args{
				str: "Hello_World",
				sep: '_',
			},
			"hello_world",
		},
		{
			"小写带分隔符",
			args{
				str: "HelloWor_ldID",
				sep: '_',
			},
			"hello_wor_ld_id",
		},
		{
			"小写带分隔符",
			args{
				str: "HelloWor_ldA",
				sep: '_',
			},
			"hello_wor_ld_a",
		},
		{
			"特殊分隔IDCom",
			args{
				str: "IDCom",
				sep: '_',
			},
			"id_com",
		},
		{
			"特殊分隔IDcom",
			args{
				str: "IDcom",
				sep: '_',
			},
			"i_dcom",
		},
		{
			"特殊分隔nameIDCom",
			args{
				str: "nameIDCom",
				sep: '_',
			},
			"name_id_com",
		},
		{
			"特殊分隔nameIDcom",
			args{
				str: "nameIDcom",
				sep: '_',
			},
			"name_i_dcom",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Recombine(tt.args.str, tt.args.sep); got != tt.want {
				t.Errorf("Recombine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnRecombine(t *testing.T) {
	type args struct {
		str string
		sep byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"空字符",
			args{
				str: "",
				sep: '_',
			},
			"",
		},
		{
			"以_分隔符",
			args{
				str: "hello_world",
				sep: '_',
			},
			"HelloWorld",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnRecombine(tt.args.str, tt.args.sep); got != tt.want {
				t.Errorf("UnRecombine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecode_Recombine(t *testing.T) {
	type args struct {
		str string
		sep byte
	}
	tests := []struct {
		name   string
		recode *Recode
		args   args
		want   string
	}{
		{
			"空字符串",
			NewRecode([]string{}),
			args{
				str: "",
				sep: '_',
			},
			"",
		},
		{
			"特殊字符IDCom",
			NewRecode([]string{"ID"}),
			args{
				str: "IDCom",
				sep: '_',
			},
			"id_com",
		},
		{
			"特殊字符IDcom",
			NewRecode([]string{"ID"}),
			args{
				str: "IDcom",
				sep: '_',
			},
			"idcom",
		},
		{
			"特殊字符nameIDCom",
			NewRecode([]string{"ID"}),
			args{
				str: "nameIDCom",
				sep: '_',
			},
			"name_id_com",
		},
		{
			"特殊字符nameIDcom",
			NewRecode([]string{"ID"}),
			args{
				str: "nameIDcom",
				sep: '_',
			},
			"name_idcom",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.recode.Recombine(tt.args.str, tt.args.sep); got != tt.want {
				t.Errorf("Recombine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnakeCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"空字符串",
			args{str: ""},
			"",
		},
		{
			"特殊字符IDCom",
			args{str: "IDCom"},
			"id_com",
		},
		{
			"特殊字符IDcom",
			args{str: "IDcom"},
			"idcom",
		},
		{
			"特殊字符nameIDCom",
			args{str: "nameIDCom"},
			"name_id_com",
		},
		{
			"特殊字符nameIDcom",
			args{str: "nameIDcom"},
			"name_idcom",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SnakeCase(tt.args.str); got != tt.want {
				t.Errorf("SnakeCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKebab(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"空字符串",
			args{str: ""},
			"",
		},
		{
			"特殊字符IDCom",
			args{str: "IDCom"},
			"id-com",
		},
		{
			"特殊字符IDcom",
			args{str: "IDcom"},
			"idcom",
		},
		{
			"特殊字符nameIDCom",
			args{str: "nameIDCom"},
			"name-id-com",
		},
		{
			"特殊字符nameIDcom",
			args{str: "nameIDcom"},
			"name-idcom",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Kebab(tt.args.str); got != tt.want {
				t.Errorf("SnakeCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCamelCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"空字符串",
			args{str: ""},
			"",
		},
		{
			"特殊字符IDCom",
			args{str: "id_com"},
			"IdCom",
		},
		{
			"特殊字符IDcom",
			args{str: "idcom"},
			"Idcom",
		},
		{
			"特殊字符nameIDCom",
			args{str: "name_id_com"},
			"NameIdCom",
		},
		{
			"特殊字符nameIDcom",
			args{str: "name_idcom"},
			"NameIdcom",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CamelCase(tt.args.str); got != tt.want {
				t.Errorf("CamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowTitle(t *testing.T) {
	var LowTitleTests = []struct {
		in, out string
	}{
		{"", ""},
		{"A", "a"},
		{" aaa aaa aaa ", " aaa aaa aaa "},
		{" Aaa Aaa Aaa ", " aaa aaa aaa "},
		{"123a456", "123a456"},
		{"Double-Blind", "double-blind"},
		{"Ÿøû", "ÿøû"},
		{"With_underscore", "with_underscore"},
		{"Unicode \xe2\x80\xa8 Line Separator", "unicode \xe2\x80\xa8 line separator"},
	}
	for _, tt := range LowTitleTests {
		if s := LowTitle(tt.in); s != tt.out {
			t.Errorf("LowTitle(%q) = %q, want %q", tt.in, s, tt.out)
		}
	}
}
