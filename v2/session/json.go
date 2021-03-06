package session

import (
	session "github.com/nspcc-dev/neofs-api-go/v2/session/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func (c *ObjectSessionContext) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		ObjectSessionContextToGRPCMessage(c),
	)
}

func (c *ObjectSessionContext) UnmarshalJSON(data []byte) error {
	msg := new(session.ObjectSessionContext)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*c = *ObjectSessionContextFromGRPCMessage(msg)

	return nil
}

func (l *TokenLifetime) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		TokenLifetimeToGRPCMessage(l),
	)
}

func (l *TokenLifetime) UnmarshalJSON(data []byte) error {
	msg := new(session.SessionToken_Body_TokenLifetime)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*l = *TokenLifetimeFromGRPCMessage(msg)

	return nil
}

func (t *SessionTokenBody) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		SessionTokenBodyToGRPCMessage(t),
	)
}

func (t *SessionTokenBody) UnmarshalJSON(data []byte) error {
	msg := new(session.SessionToken_Body)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*t = *SessionTokenBodyFromGRPCMessage(msg)

	return nil
}

func (t *SessionToken) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		SessionTokenToGRPCMessage(t),
	)
}

func (t *SessionToken) UnmarshalJSON(data []byte) error {
	msg := new(session.SessionToken)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*t = *SessionTokenFromGRPCMessage(msg)

	return nil
}

func (x *XHeader) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		XHeaderToGRPCMessage(x),
	)
}

func (x *XHeader) UnmarshalJSON(data []byte) error {
	msg := new(session.XHeader)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*x = *XHeaderFromGRPCMessage(msg)

	return nil
}

func (r *RequestMetaHeader) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		RequestMetaHeaderToGRPCMessage(r),
	)
}

func (r *RequestMetaHeader) UnmarshalJSON(data []byte) error {
	msg := new(session.RequestMetaHeader)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*r = *RequestMetaHeaderFromGRPCMessage(msg)

	return nil
}

func (r *RequestVerificationHeader) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		RequestVerificationHeaderToGRPCMessage(r),
	)
}

func (r *RequestVerificationHeader) UnmarshalJSON(data []byte) error {
	msg := new(session.RequestVerificationHeader)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*r = *RequestVerificationHeaderFromGRPCMessage(msg)

	return nil
}

func (r *ResponseMetaHeader) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		ResponseMetaHeaderToGRPCMessage(r),
	)
}

func (r *ResponseMetaHeader) UnmarshalJSON(data []byte) error {
	msg := new(session.ResponseMetaHeader)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*r = *ResponseMetaHeaderFromGRPCMessage(msg)

	return nil
}

func (r *ResponseVerificationHeader) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(
		ResponseVerificationHeaderToGRPCMessage(r),
	)
}

func (r *ResponseVerificationHeader) UnmarshalJSON(data []byte) error {
	msg := new(session.ResponseVerificationHeader)

	if err := protojson.Unmarshal(data, msg); err != nil {
		return err
	}

	*r = *ResponseVerificationHeaderFromGRPCMessage(msg)

	return nil
}
