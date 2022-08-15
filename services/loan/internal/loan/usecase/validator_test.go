package usecase

import (
	"context"
	"testing"
)

func Test_validateName(t *testing.T) {
	ctx := context.Background()

	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Name only contain 1 word",
			args: args{
				ctx:  ctx,
				name: "alfin",
			},
			want: false,
		},
		{
			name: "Name contain 2 word",
			args: args{
				ctx:  ctx,
				name: "alfin lie",
			},
			want: true,
		},
		{
			name: "Name contain more than 2 word",
			args: args{
				ctx:  ctx,
				name: "alfin christian lie",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateName(tt.args.ctx, tt.args.name); got != tt.want {
				t.Errorf("validateName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateIdentityNumber(t *testing.T) {
	ctx := context.Background()

	type args struct {
		ctx            context.Context
		sex            string
		dateOfBirth    string
		identityNumber string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Error parse date of birth",
			args: args{
				ctx:            ctx,
				sex:            "MALE",
				dateOfBirth:    "123123",
				identityNumber: "1234560101954567",
			},
			want: false,
		},
		{
			name: "Error male identity number",
			args: args{
				ctx:            ctx,
				sex:            "MALE",
				dateOfBirth:    "01-01-1995",
				identityNumber: "1234560501954567",
			},
			want: false,
		},
		{
			name: "Success male identity number",
			args: args{
				ctx:            ctx,
				sex:            "MALE",
				dateOfBirth:    "01-01-1995",
				identityNumber: "1234560101954567",
			},
			want: true,
		},
		{
			name: "Error female identity number",
			args: args{
				ctx:            ctx,
				sex:            "FEMALE",
				dateOfBirth:    "01-01-1995",
				identityNumber: "1234560501954567",
			},
			want: false,
		},
		{
			name: "Success female identity number",
			args: args{
				ctx:            ctx,
				sex:            "FEMALE",
				dateOfBirth:    "01-01-1995",
				identityNumber: "1234564101954567",
			},
			want: true,
		},
		{
			name: "Error other gender identity number",
			args: args{
				ctx:            ctx,
				sex:            "OTHER",
				dateOfBirth:    "01-01-1995",
				identityNumber: "1234560101954567",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateIdentityNumber(tt.args.ctx, tt.args.sex, tt.args.dateOfBirth, tt.args.identityNumber); got != tt.want {
				t.Errorf("validateIdentityNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateAmount(t *testing.T) {
	ctx := context.Background()

	type args struct {
		ctx    context.Context
		amount int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Under 1 mio",
			args: args{
				ctx:    ctx,
				amount: 500000,
			},
			want: false,
		},
		{
			name: "Between 1 mio and 10 mio",
			args: args{
				ctx:    ctx,
				amount: 5000000,
			},
			want: true,
		},
		{
			name: "Over 10 mio",
			args: args{
				ctx:    ctx,
				amount: 50000000,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateAmount(tt.args.ctx, tt.args.amount); got != tt.want {
				t.Errorf("validateAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validatePurpose(t *testing.T) {
	ctx := context.Background()

	type args struct {
		ctx     context.Context
		purpose string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Contain 0 specific words",
			args: args{
				ctx:     ctx,
				purpose: "Buy anything",
			},
			want: false,
		},
		{
			name: "Contain 1 specific words",
			args: args{
				ctx:     ctx,
				purpose: "Buy car",
			},
			want: true,
		},
		{
			name: "Contain more than 1 specific words",
			args: args{
				ctx:     ctx,
				purpose: "Buy car and electronics",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validatePurpose(tt.args.ctx, tt.args.purpose); got != tt.want {
				t.Errorf("validatePurpose() = %v, want %v", got, tt.want)
			}
		})
	}
}
