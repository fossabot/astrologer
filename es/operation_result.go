package es

import (
	"github.com/stellar/go/xdr"
)

// AppendResult Appends operation result
func AppendResult(op *Operation, r *xdr.OperationResult) {
	op.ResultCode = int(r.Code)

	if r.Code == xdr.OperationResultCodeOpInner {
		switch t := r.Tr.Type; t {
		case xdr.OperationTypeCreateAccount:
			newCreateAccountResult(r.Tr.MustCreateAccountResult(), op)
		case xdr.OperationTypePayment:
			newPaymentResult(r.Tr.MustPaymentResult(), op)
		case xdr.OperationTypePathPayment:
			newPathPaymentResult(r.Tr.MustPathPaymentResult(), op)
		case xdr.OperationTypeManageOffer:
			newManageOfferResult(r.Tr.MustManageOfferResult(), op)
		case xdr.OperationTypeCreatePassiveOffer:
			newManageOfferResult(r.Tr.MustCreatePassiveOfferResult(), op)
		case xdr.OperationTypeSetOptions:
			newSetOptionsResult(r.Tr.MustSetOptionsResult(), op)
		case xdr.OperationTypeChangeTrust:
			newChangeTrustResult(r.Tr.MustChangeTrustResult(), op)
		case xdr.OperationTypeAllowTrust:
			newAllowTrustResult(r.Tr.MustAllowTrustResult(), op)
		case xdr.OperationTypeAccountMerge:
			newAccountMergeResult(r.Tr.MustAccountMergeResult(), op)
		case xdr.OperationTypeManageData:
			newManageDataResult(r.Tr.MustManageDataResult(), op)
		case xdr.OperationTypeBumpSequence:
			newBumpSequenceResult(r.Tr.MustBumpSeqResult(), op)
		}
	} else {
		op.Succesful = false
	}
}

func newCreateAccountResult(r xdr.CreateAccountResult, op *Operation) {
	op.InnerResultCode = int(r.Code)
	op.Succesful = r.Code == xdr.CreateAccountResultCodeCreateAccountSuccess
}

func newPaymentResult(r xdr.PaymentResult, op *Operation) {
	op.InnerResultCode = int(r.Code)
	op.Succesful = r.Code == xdr.PaymentResultCodePaymentSuccess
}

func newPathPaymentResult(r xdr.PathPaymentResult, op *Operation) {
	op.InnerResultCode = int(r.Code)
	op.Succesful = r.Code == xdr.PathPaymentResultCodePathPaymentSuccess
}

func newManageOfferResult(r xdr.ManageOfferResult, op *Operation) {
	op.InnerResultCode = int(r.Code)
	op.Succesful = r.Code == xdr.ManageOfferResultCodeManageOfferSuccess
}

func newSetOptionsResult(r xdr.SetOptionsResult, op *Operation) {
	op.InnerResultCode = int(r.Code)
	op.Succesful = r.Code == xdr.SetOptionsResultCodeSetOptionsSuccess
}

func newChangeTrustResult(r xdr.ChangeTrustResult, op *Operation) {
	op.InnerResultCode = int(r.Code)
	op.Succesful = r.Code == xdr.ChangeTrustResultCodeChangeTrustSuccess
}

func newAllowTrustResult(r xdr.AllowTrustResult, op *Operation) {
	op.InnerResultCode = int(r.Code)
	op.Succesful = r.Code == xdr.AllowTrustResultCodeAllowTrustSuccess
}

func newAccountMergeResult(r xdr.AccountMergeResult, op *Operation) {
	op.InnerResultCode = int(r.Code)
	op.Succesful = r.Code == xdr.AccountMergeResultCodeAccountMergeSuccess
}

func newManageDataResult(r xdr.ManageDataResult, op *Operation) {
	op.InnerResultCode = int(r.Code)
	op.Succesful = r.Code == xdr.ManageDataResultCodeManageDataSuccess
}

func newBumpSequenceResult(r xdr.BumpSequenceResult, op *Operation) {
	op.InnerResultCode = int(r.Code)
	op.Succesful = r.Code == xdr.BumpSequenceResultCodeBumpSequenceSuccess
}
