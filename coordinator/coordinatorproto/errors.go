package coordinatorproto

import (
	"errors"

	"github.com/anyproto/any-sync/net/rpc/rpcerr"
)

var (
	errGroup = rpcerr.ErrGroup(ErrorCodes_ErrorOffset)

	ErrUnexpected           = errGroup.Register(errors.New("unexpected error"), uint64(ErrorCodes_Unexpected))
	ErrSpaceIsCreated       = errGroup.Register(errors.New("space is missing"), uint64(ErrorCodes_SpaceCreated))
	ErrSpaceIsDeleted       = errGroup.Register(errors.New("space is deleted"), uint64(ErrorCodes_SpaceDeleted))
	ErrSpaceDeletionPending = errGroup.Register(errors.New("space is set out for deletion"), uint64(ErrorCodes_SpaceDeletionPending))
	ErrSpaceNotExists       = errGroup.Register(errors.New("space not exists"), uint64(ErrorCodes_SpaceNotExists))
	ErrSpaceLimitReached    = errGroup.Register(errors.New("space limit reached"), uint64(ErrorCodes_SpaceLimitReached))
	ErrAccountIsDeleted     = errGroup.Register(errors.New("account is deleted"), uint64(ErrorCodes_AccountDeleted))
)
