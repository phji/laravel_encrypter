package encrypter

import (
	"reflect"
	"testing"
)

func TestEncrypter_hash(t *testing.T) {
	type fields struct {
		key string
	}
	type args struct {
		iv    string
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			fields: fields{key: "qyk5OUGEoI3e7asY/ij+uMEeBnSxWTDS8LT7ExX1u88="},
			args:   args{iv: "", value: ""},
			want:   "794de6a7f3806fad54729599045b14a7b854702bbdd931c1deeff085e45b8d03",
		},
		{
			fields: fields{key: "qyk5OUGEoI3e7asY/ij+uMEeBnSxWTDS8LT7ExX1u88="},
			args:   args{iv: "aaa", value: ""},
			want:   "80a3ecfd25176b3a4e8a31a548d170ebab47fd2edd263c1c25ae5090b1604986",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NewEncrypter(tt.fields.key)

			if got := e.hash(tt.args.iv, tt.args.value); got != tt.want {
				t.Errorf("hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncrypter_getJsonPayload(t *testing.T) {
	type fields struct {
		key string
	}
	type args struct {
		payload string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Payload
	}{
		{
			fields: fields{key: "qyk5OUGEoI3e7asY/ij+uMEeBnSxWTDS8LT7ExX1u88="},
			args:   args{payload: "eyJpdiI6IlliTVRuZXZVcDEyWlhRRWp4Vjd3TWc9PSIsInZhbHVlIjoiMTdCbkdvR3lOanNaOWxzMWZtYkJuUT09IiwibWFjIjoiMzJmODU1NmY1NDZkZDFlZTJlZjE2M2ZiOWNiODY2NDRlMTY5YTRhYTVlNmIxN2JjZWU1MGIzZTc1OWViZmQyNSJ9"},
			want:   Payload{Iv: "YbMTnevUp12ZXQEjxV7wMg==", Value: "17BnGoGyNjsZ9ls1fmbBnQ==", Mac: "32f8556f546dd1ee2ef163fb9cb86644e169a4aa5e6b17bcee50b3e759ebfd25"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NewEncrypter(tt.fields.key)
			if got := e.getJsonPayload(tt.args.payload); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getJsonPayload() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncrypter_Decrypt(t *testing.T) {
	type fields struct {
		key string
	}
	type args struct {
		payload string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		wantErr bool
	}{
		{
			name: "a",
			fields: fields{key: "qyk5OUGEoI3e7asY/ij+uMEeBnSxWTDS8LT7ExX1u88="},
			args:   args{payload: "eyJpdiI6IlliTVRuZXZVcDEyWlhRRWp4Vjd3TWc9PSIsInZhbHVlIjoiMTdCbkdvR3lOanNaOWxzMWZtYkJuUT09IiwibWFjIjoiMzJmODU1NmY1NDZkZDFlZTJlZjE2M2ZiOWNiODY2NDRlMTY5YTRhYTVlNmIxN2JjZWU1MGIzZTc1OWViZmQyNSJ9"},
			want:   "a",
			wantErr: false,
		},
		{
			name: "empty",
			fields: fields{key: "qyk5OUGEoI3e7asY/ij+uMEeBnSxWTDS8LT7ExX1u88="},
			args:   args{payload: "eyJpdiI6InR3Z3pkQlJTNmQzMzJpUXo4ME5EZ0E9PSIsInZhbHVlIjoiYjVHY095alN2Ym1ZOEZWaUkyZW9tQT09IiwibWFjIjoiODIwMjE3NjQwOTZmMzM2MTgzMDIwYTY0NGQwOWI4NmRiNjNmN2Q4MjliMDg1NWQ4OWVkZTQwZDgzZjg2MzU2ZCJ9"},
			want:   "",
			wantErr: false,
		},
		{
			name: "16 byte",
			fields: fields{key: "qyk5OUGEoI3e7asY/ij+uMEeBnSxWTDS8LT7ExX1u88="},
			args:   args{payload: "eyJpdiI6Im93bHZkNXlpdXg5Vmg2M2YxNW1tbnc9PSIsInZhbHVlIjoiUThMUlBvd21VUkFFV0ZzUCt0MmUyWXgvSjVvVmtuMDRQdVFmMlc2R0lNVT0iLCJtYWMiOiI5MzM2NTFhMDg2ZTZhZmI4OTM2ODk2NThjOTQ3OWY4NGEwNmExNTFkNzg3NTBiZGQ2NjgxYjFiYjc3MzE1OWMwIn0="},
			want:   "aaaaaaaaaaaaaaaa",
			wantErr: false,
		},
		{
			name: "iv error",
			fields: fields{key: "qyk5OUGEoI3e7asY/ij+uMEeBnSxWTDS8LT7ExX1u88="},
			args:   args{payload: "eyJpdiI6ImEiLCJ2YWx1ZSI6IiIsIm1hYyI6IiJ9"},
			wantErr: true,
		},
		{
			name: "len(cipherText) < aes.BlockSize",
			fields: fields{key: "qyk5OUGEoI3e7asY/ij+uMEeBnSxWTDS8LT7ExX1u88="},
			args:   args{payload: "eyJpdiI6IiIsInZhbHVlIjoiIiwibWFjIjoiIn0="},
			wantErr: true,
		},
		{
			name: "len(cipherText)%aes.BlockSize != 0",
			fields: fields{key: "qyk5OUGEoI3e7asY/ij+uMEeBnSxWTDS8LT7ExX1u88="},
			args:   args{payload: "eyJpdiI6ImFhYWFhYWFhYWFhYWFhYWEiLCJ2YWx1ZSI6IiIsIm1hYyI6IiJ9"},
			wantErr: true,
		},
		{
			name: "key error",
			fields: fields{key: ""},
			args:   args{payload: "eyJpdiI6Im93bHZkNXlpdXg5Vmg2M2YxNW1tbnc9PSIsInZhbHVlIjoiUThMUlBvd21VUkFFV0ZzUCt0MmUyWXgvSjVvVmtuMDRQdVFmMlc2R0lNVT0iLCJtYWMiOiI5MzM2NTFhMDg2ZTZhZmI4OTM2ODk2NThjOTQ3OWY4NGEwNmExNTFkNzg3NTBiZGQ2NjgxYjFiYjc3MzE1OWMwIn0="},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NewEncrypter(tt.fields.key)
			got, err := e.Decrypt(tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncrypter_Encrypt(t *testing.T) {
	type fields struct {
		key string
	}
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		wantErr bool
	}{
		{
			name: "empty encrypt -> decrypt want empty",
			fields: fields{key: "qyk5OUGEoI3e7asY/ij+uMEeBnSxWTDS8LT7ExX1u88="},
			args: args{value: ""},
			//  eyJpdiI6IlNnakF1SmNETXZrRG13cjYyM2cyOFE9PSIsInZhbHVlIjoiVDVLbEhCTUYrVTVMcnNITVAzYzVIdz09IiwibWFjIjoiMTMxNTJkOGFiNGI5NDZjNzEyNmYyODJlMGM4ZjM3ODI4OTA3NmE4OGFiZGZmMmQxYzRjNTdjMDQyZDZmZGNmZSJ9
			want: "",
			wantErr: false,
		},
		{
			name: "a encrypt -> decrypt want a",
			fields: fields{key: "qyk5OUGEoI3e7asY/ij+uMEeBnSxWTDS8LT7ExX1u88="},
			args: args{value: "a"},
			want: "a",
			wantErr: false,
		},
		{
			name: "16 byte",
			fields: fields{key: "qyk5OUGEoI3e7asY/ij+uMEeBnSxWTDS8LT7ExX1u88="},
			args: args{value: "aaaaaaaaaaaaaaaa"},
			want: "aaaaaaaaaaaaaaaa",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NewEncrypter(tt.fields.key)

			if got, _ := e.Decrypt(e.Encrypt(tt.args.value)); got != tt.want {
				t.Errorf("Encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncrypter_setJsonPayload(t *testing.T) {
	type fields struct {
		key string
	}
	type args struct {
		payload Payload
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			fields: fields{key: "qyk5OUGEoI3e7asY/ij+uMEeBnSxWTDS8LT7ExX1u88="},
			args: args{payload: Payload{Iv: "", Value: "", Mac: ""}},
			want: "eyJpdiI6IiIsInZhbHVlIjoiIiwibWFjIjoiIn0=",
			wantErr: false,
		},
		{
			fields: fields{key: "qyk5OUGEoI3e7asY/ij+uMEeBnSxWTDS8LT7ExX1u88="},
			args: args{payload: Payload{Iv: "a", Value: "", Mac: ""}},
			want: "eyJpdiI6ImEiLCJ2YWx1ZSI6IiIsIm1hYyI6IiJ9",
			wantErr: false,
		},
		{
			fields: fields{key: "qyk5OUGEoI3e7asY/ij+uMEeBnSxWTDS8LT7ExX1u88="},
			args: args{payload: Payload{Iv: "aaaaaaaaaaaaaaaa", Value: "", Mac: ""}},
			want: "eyJpdiI6ImFhYWFhYWFhYWFhYWFhYWEiLCJ2YWx1ZSI6IiIsIm1hYyI6IiJ9",
			wantErr: false,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NewEncrypter(tt.fields.key)
			got, err := e.setJsonPayload(tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("setJsonPayload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("setJsonPayload() got = %v, want %v", got, tt.want)
			}
		})
	}
}