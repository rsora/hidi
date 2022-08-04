package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestScrambleAWSResourceID(t *testing.T) {

	tests := []struct {
		name         string
		line         string
		expectedLine string
	}{
		{
			name:         "EC2 instance-id short version",
			line:         "i-b9b4ffaa",
			expectedLine: "i-8ee7a961",
		},
		{
			name:         "EC2 instance-id long version",
			line:         "i-ccd8f9b742c3030e1",
			expectedLine: "i-669e0fa0af90372f2",
		},
		{
			name:         "EC2 volume-id short version",
			line:         "vol-b9b4ffaa",
			expectedLine: "vol-8ee7a961",
		},
		{
			name:         "EC2 volume-id long version",
			line:         "vol-b365e5cb5e48fe90f",
			expectedLine: "vol-1cfa6ca7ea18c6f94",
		},
		{
			name:         "EC2 snapshot-id long version",
			line:         "snap-5a4bbca41abf53995",
			expectedLine: "snap-b93053e63edda104f",
		},
		{
			name:         "EC2 image-id (AMI id) short version",
			line:         "ami-ada048d9",
			expectedLine: "ami-1083a7b5",
		},
		{
			name:         "VPC subnet-id short version",
			line:         "subnet-dcdf41c6",
			expectedLine: "subnet-3851c2eb",
		},
		{
			name:         "VPC vpc-id short version",
			line:         "vpc-7fe5842f",
			expectedLine: "vpc-d99aecba",
		},
		{
			name:         "EC2 ENI attachment-id long version",
			line:         "eni-attach-3968b50acb38d32df",
			expectedLine: "eni-attach-d2d756ffde2306465",
		},
		{
			name:         "EC2 ENI interface-id long version",
			line:         "eni-73d19acaee09d272e",
			expectedLine: "eni-b9c45855dfd6bf468",
		},
		{
			name:         "EC2 Security Group id long version",
			line:         "sg-9a4c9f0ce29230f1b",
			expectedLine: "sg-3d149d1fcd7ab7e2e",
		},
		{
			name:         "EC2 Reservation reservation-id long version",
			line:         "r-efc725906c6e425fc",
			expectedLine: "r-105eddefcfca20cc4",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scrambledLine := ScrambleAWSResourceID(tt.line, "salt")
			if !cmp.Equal(scrambledLine, tt.expectedLine) {
				t.Errorf("Wrong line generated, diff:\n%s", cmp.Diff(tt.expectedLine, scrambledLine))
			}
		})
	}
}

func TestScrambleAWSResourceIDInARealLine(t *testing.T) {

	tests := []struct {
		name         string
		line         string
		expectedLine string
	}{
		{
			name:         "EC2 instance-id replace in a random placebo response file line",
			line:         `                        "InstanceId": "i-ccd8f9b742c3030e1",`,
			expectedLine: `                        "InstanceId": "i-669e0fa0af90372f2",`,
		},
		{
			name:         "do not replace anything this string does not contain any aws resource id",
			line:         `                "Name": "emr-01-dev-eks-janitor-20180510010042153206",`,
			expectedLine: `                "Name": "emr-01-dev-eks-janitor-20180510010042153206",`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scrambledLine := ScrambleAWSResourceID(tt.line, "salt")
			if !cmp.Equal(scrambledLine, tt.expectedLine) {
				t.Errorf("Wrong line generated, diff:\n%s", cmp.Diff(tt.expectedLine, scrambledLine))
			}
		})
	}
}
