package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"mindfit.noti/graph/generated"
	"mindfit.noti/graph/model"

	///Users/tong/gitlab/minefit_noti_demo_07-21/system/notification
	noti "mindfit.noti/system/notification"
)

func (r *mutationResolver) CreateNoti(ctx context.Context, input model.NewNoti) (string, error) {
	var userid_in noti.Idonly
	userid_in.Userid = input.Userid
	var resultNotiSchema []*model.NotiSchema = noti.Getnoti(userid_in)

	resultNotiSchema = append(resultNotiSchema, &model.NotiSchema{Name: input.Name, Related: input.Related, Type: input.Type, Appointment: input.Appointment})
	fmt.Println(resultNotiSchema)

	var updatedata noti.ReturnNotiFromDB
	updatedata.Userid = input.Userid
	updatedata.Notilist = resultNotiSchema
	var out bool = noti.Updatenoti(updatedata)

	if out {
		fmt.Println("update")
	}

	return "OK", nil
}

func (r *mutationResolver) CallNoti(ctx context.Context, input model.UserWhoCall) (*model.ReturnNoti, error) {
	var userid_in noti.Idonly
	userid_in.Userid = input.Userid
	var out *model.ReturnNoti
	var resultNotiSchema []*model.NotiSchema = noti.Getnoti(userid_in)
	fmt.Println(resultNotiSchema)
	// resultNotiSchema = append(resultNotiSchema, &model.NotiSchema{Name: "Dr.Parin", Related: "Doctor", Type: "", Appointment: "15-07-21 | 09:00-10:00"})
	// resultNotiSchema = append(resultNotiSchema, &model.NotiSchema{Name: "Tho-ng", Related: "Personal", Type: "Teletherapy", Doctor: "Dr.Parin", Appointment: "15-07-21 | 10:00-11:00"})

	dummyOut := model.ReturnNoti{
		Notilist: resultNotiSchema,
	}
	out = &dummyOut
	return out, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
