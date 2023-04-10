package schema

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestCreateUserIDEmail(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer UserIDEmail_c.Drop()

	expected0 := &UserIDEmail{
		UserID:  "SYSOP",
		IDEmail: "test@ptt.test",

		UpdateNanoTS: 1234567890000000000,
	}

	expected1 := &UserIDEmail{
		UserID:  "SYSOP2",
		IDEmail: "test@ptt2.test",

		UpdateNanoTS: 1234567890000000000,
	}

	expected2 := &UserIDEmail{
		UserID:  "SYSOP",
		IDEmail: "test@ptt3.test",

		UpdateNanoTS: 1234589890000000000,
	}

	errUnique := mongo.WriteException{
		WriteErrors: mongo.WriteErrors([]mongo.WriteError{
			{Code: 11000, Message: "E11000 duplicate key error collection: devptt_test.user_id_email index: idemail_1 dup key: { idemail: \"test@ptt2.test\" }"},
		}),
	}
	type args struct {
		userID       bbs.UUserID
		email        string
		updateNanoTS types.NanoTS
	}
	tests := []struct {
		name             string
		args             args
		wantErr          bool
		expectedErr      error
		expected         *UserIDEmail
		expectedByUserID *UserIDEmail
	}{
		// TODO: Add test cases.
		{
			name:             "SYSOP-ptt",
			args:             args{userID: "SYSOP", email: "test@ptt.test", updateNanoTS: 1234567890000000000},
			expected:         expected0,
			expectedByUserID: expected0,
		},
		{
			name:             "SYSOP2-ptt2",
			args:             args{userID: "SYSOP2", email: "test@ptt2.test", updateNanoTS: 1234567890000000000},
			expected:         expected1,
			expectedByUserID: expected1,
		},
		{
			name:             "SYSOP-ptt2: no-create",
			args:             args{userID: "SYSOP", email: "test@ptt2.test", updateNanoTS: 1234567890000000000},
			wantErr:          true,
			expectedErr:      ErrNoCreate,
			expected:         expected1,
			expectedByUserID: expected0,
		},
		{
			name:             "SYSOP-ptt3-not-expired",
			args:             args{userID: "SYSOP", email: "test@ptt3.test", updateNanoTS: 1234567890000000000},
			wantErr:          true,
			expectedErr:      ErrNoCreate,
			expected:         nil,
			expectedByUserID: expected0,
		},
		{
			name:             "SYSOP-ptt2-expired, not unique",
			args:             args{userID: "SYSOP", email: "test@ptt2.test", updateNanoTS: 1234587890000000000},
			wantErr:          true,
			expectedErr:      errUnique,
			expected:         nil,
			expectedByUserID: nil,
		},
		{
			name:             "SYSOP-ptt3: expired",
			args:             args{userID: "SYSOP", email: "test@ptt3.test", updateNanoTS: 1234589890000000000},
			expected:         expected2,
			expectedByUserID: expected2,
		},
	}

	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			err := CreateUserIDEmail(tt.args.userID, tt.args.email, tt.args.updateNanoTS)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUserIDEmail() error = %v, wantErr %v", err, tt.wantErr)
			}

			switch err.(type) {
			case mongo.WriteException:
				err2, ok := err.(mongo.WriteException)
				if ok {
					err2.Raw = nil
					writeErrors := make([]mongo.WriteError, 0, len(err2.WriteErrors))
					for _, each := range err2.WriteErrors {
						each.Raw = nil
						writeErrors = append(writeErrors, each)
					}
					err2.WriteErrors = writeErrors
					assert.Equal(t, tt.expectedErr, err2)
				}
			default:
				assert.Equal(t, tt.expectedErr, err)
			}

			got, _ := GetUserIDEmailByEmail(tt.args.email, tt.args.updateNanoTS)
			testutil.TDeepEqual(t, "got", got, tt.expected)

			got, _ = GetUserIDEmailByUserID(tt.args.userID, tt.args.updateNanoTS)
			testutil.TDeepEqual(t, "gotByUserID", got, tt.expectedByUserID)
		})
		wg.Wait()
	}
}

func TestUpdateUserIDEmailIsSet(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer UserIDEmail_c.Drop()

	_ = CreateUserIDEmail("SYSOP", "test@ptt.test", 1234567890000000)

	type args struct {
		userID       bbs.UUserID
		email        string
		isSet        bool
		updateNanoTS types.NanoTS
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{userID: "SYSOP", email: "test@ptt.test", isSet: true, updateNanoTS: 1234567890000000001},
		},
		{
			args:    args{userID: "SYSOP", email: "test@ptt2.test", isSet: true, updateNanoTS: 1234567890000000002},
			wantErr: true,
		},
		{
			args:    args{userID: "SYSOP2", email: "test@ptt.test", isSet: true, updateNanoTS: 1234567890000000003},
			wantErr: true,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if err := UpdateUserIDEmailIsSet(tt.args.userID, tt.args.email, tt.args.isSet, tt.args.updateNanoTS); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUserIDEmailIsSet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		wg.Wait()
	}
}
