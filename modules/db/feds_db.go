package db

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func removeInt(s bson.A, r int64) bson.A {
	for i, v := range s {
		if v.(int64) == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func stringInSliceA(a string, list bson.A) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func deduplicate_fban(s bson.A, x int64) (bson.A, bool, int) {
	for i, v := range s {
		if v.(bson.M)["user_id"].(int64) == x {
			return append(s[:i], s[i+1:]...), true, i
		}
	}
	return s, false, 0
}

var feds = database.Collection("fed1")
var fed_chats = database.Collection("fedch1")
var fedadmins = database.Collection("fedadmi1")
var fbans = database.Collection("fbani1")

func Make_new_fed(user_id int64, fedname string) (string, string) {
	uid := uuid.New().String()
	filter := bson.M{"user_id": user_id}
	feds.UpdateOne(context.TODO(), filter, bson.D{{Key: "$set", Value: bson.D{{Key: "fed_id", Value: uid}, {Key: "fedname", Value: fedname}, {Key: "flog", Value: nil}, {Key: "chats", Value: []int64{}}, {Key: "report", Value: true}, {Key: "fadmins", Value: []int64{}}}}}, opts)
	return uid, fedname
}

func Get_fed_by_owner(user_id int64) (bool, string, string) {
	filter := bson.M{"user_id": user_id}
	fed := feds.FindOne(context.TODO(), filter)
	if fed.Err() != nil {
		return false, "", ""
	}
	var fed_info bson.M
	fed.Decode(&fed_info)
	return true, fed_info["fed_id"].(string), fed_info["fedname"].(string)
}

func Delete_fed(fed_id string) {
	filter := bson.M{"fed_id": fed_id}
	feds.DeleteOne(context.TODO(), filter)
}

func Rename_fed_by_id(fed_id string, name string) {
	filter := bson.M{"fed_id": fed_id}
	feds.UpdateOne(context.TODO(), filter, bson.D{{Key: "$set", Value: bson.D{{Key: "fedname", Value: name}}}}, opts)
}

func Transfer_fed(fed_id string, user_id int64) {
	feds.UpdateOne(context.TODO(), bson.M{"fed_id": fed_id}, bson.D{{Key: "$set", Value: bson.D{{Key: "user_id", Value: user_id}}}}, opts)
}

func Chat_join_fed(fed_id string, chat_id int64) {
        chats := bson.A{chat_id}
	filter := bson.M{"chat_id": chat_id}
	f := fed_chats.FindOne(context.TODO(), filter)
	if f.Err() == nil {
		fed_chats.DeleteOne(context.TODO(), filter)
		var chats_m bson.M
		feds.FindOne(context.TODO(), bson.M{"fed_id": fed_id}).Decode(&chats_m)
		feds.UpdateOne(context.TODO(), bson.M{"fed_id": fed_id}, bson.D{{Key: "$set", Value: bson.D{{Key: "chats", Value: removeInt(chats_m["chats"].(bson.A), chat_id)}}}}, opts)
	}
	fed_chats.UpdateOne(context.TODO(), filter, bson.D{{Key: "$set", Value: bson.D{{Key: "fed_id", Value: fed_id}}}}, opts)
        F := feds.FindOne(context.TODO(), bson.M{"fed_id": fed_id})
        if F.Err() == nil{
	 var chats_m bson.M
	 F.Decode(&chats_m)
         chats = append(chats_m["chats"].(bson.A), chat_id)
        }
	feds.UpdateOne(context.TODO(), bson.M{"fed_id": fed_id}, bson.D{{Key: "$set", Value: bson.D{{Key: "chats", Value: chats}}}}, opts)
}

func Chat_leave_fed(chat_id int64) {
	var chats_a bson.M
	fed_chats.FindOne(context.TODO(), bson.M{"chat_id": chat_id}).Decode(&chats_a)
	fed_id := chats_a["fed_id"].(string)
	fed_chats.DeleteOne(context.TODO(), bson.M{"chat_id": chat_id})
	var chats_m bson.M
	feds.FindOne(context.TODO(), bson.M{"fed_id": fed_id}).Decode(&chats_m)
	feds.UpdateOne(context.TODO(), bson.M{"fed_id": fed_id}, bson.D{{Key: "$set", Value: bson.D{{Key: "chats", Value: removeInt(chats_m["chats"].(bson.A), chat_id)}}}}, opts)
}

func Get_chat_fed(chat_id int64) string {
	var chats_a bson.M
	f := fed_chats.FindOne(context.TODO(), bson.M{"chat_id": chat_id})
	if f.Err() != nil {
		return ""
	} else {
		f.Decode(&chats_a)
		return chats_a["fed_id"].(string)
	}
}

func User_join_fed(fed_id string, user_id int64) bool {
	feds_list := bson.A{fed_id}
	fadmins := bson.A{user_id}
	filter := bson.M{"fed_id": fed_id}
	f := feds.FindOne(context.TODO(), filter)
	if f.Err() == nil {
		var fad bson.M
		f.Decode(&fad)
		fadmins = fad["fadmins"].(bson.A)
		fadmins = append(fadmins, user_id)
	}
	a := fedadmins.FindOne(context.TODO(), bson.M{"user_id": user_id})
	if a.Err() == nil {
		var feds_ll bson.M
		a.Decode(&feds_ll)
		feds_list = feds_ll["feds"].(bson.A)
		feds_list = append(feds_list, fed_id)
		if len(feds_list) > 20 {
			return false
		}
	}
	fedadmins.UpdateOne(context.TODO(), bson.M{"user_id": user_id}, bson.D{{Key: "$set", Value: bson.D{{Key: "feds", Value: feds_list}}}}, opts)
	feds.UpdateOne(context.TODO(), bson.M{"fed_id": fed_id}, bson.D{{Key: "$set", Value: bson.D{{Key: "fadmins", Value: fadmins}}}}, opts)
	return true
}

func User_leave_fed(fed_id string, user_id int64) {
	f := feds.FindOne(context.TODO(), bson.M{"fed_id": fed_id})
	a := fedadmins.FindOne(context.TODO(), bson.M{"user_id": user_id})
	if f.Err() == nil {
		var fg bson.M
		f.Decode(&fg)
		fadmins := removeInt(fg["fadmins"].(bson.A), user_id)
		feds.UpdateOne(context.TODO(), bson.M{"fed_id": fed_id}, bson.D{{Key: "$set", Value: bson.D{{Key: "fadmins", Value: fadmins}}}}, opts)
	}
	if a.Err() == nil {
		var fg bson.M
		a.Decode(&fg)
		feds_up := remove(fg["feds"].(bson.A), fed_id)
		feds.UpdateOne(context.TODO(), bson.M{"user_id": user_id}, bson.D{{Key: "$set", Value: bson.D{{Key: "fadmins", Value: feds_up}}}}, opts)
	}
}

func Is_user_fed_admin(user_id int64, fed_id string) bool {
	a := fedadmins.FindOne(context.TODO(), bson.M{"user_id": user_id})
	if a.Err() != nil {
		return false
	} else {
		var fg bson.M
		a.Decode(&fg)
		if stringInSliceA(fed_id, fg["feds"].(bson.A)) {
			return true
		}
	}
	return false
}

func Fban_user(user_id int64, fed_id string, reason string, name string, time_delta int64, banner int64) bool {
	var fbanned bson.A
	already_fbanned := false
	ff := fbans.FindOne(context.TODO(), bson.M{"fed_id": fed_id})
	if ff.Err() != nil {
		fbanned = append(fbanned, bson.M{"user_id": user_id, "reason": reason, "name": name, "time": time_delta, "banner": banner})
	} else {
		var document bson.M
		ff.Decode(&document)
		fbanned = document["fbans"].(bson.A)
		fb, IsBanned, _ := deduplicate_fban(fbanned, user_id)
		fbanned = fb
		if IsBanned {
			already_fbanned = true
		}
		fbanned = append(fbanned, bson.M{"user_id": user_id, "reason": reason, "name": name, "time": time_delta, "banner": banner})
	}
	fbans.UpdateOne(context.TODO(), bson.M{"fed_id": fed_id}, bson.D{{Key: "$set", Value: bson.D{{Key: "fbans", Value: fbanned}}}}, opts)
	return already_fbanned
}

func Unfban_user(user_id int64, fed_id string) bool {
	var fbanned bson.A
	ff := fbans.FindOne(context.TODO(), bson.M{"fed_id": fed_id})
	if ff.Err() != nil {
		return false
	} else {
		var document bson.M
		ff.Decode(&document)
		fbanned = document["fbans"].(bson.A)
		fbanned, IsBanned, _ := deduplicate_fban(fbanned, user_id)
		if !IsBanned {
			return false
		}
		fbans.UpdateOne(context.TODO(), bson.M{"fed_id": fed_id}, bson.D{{Key: "$set", Value: bson.D{{Key: "fbans", Value: fbanned}}}}, opts)
		return IsBanned
	}
}

func Is_Fbanned(user_id int64, fed_id string) (bool, string) {
	var fbanned bson.A
	ff := fbans.FindOne(context.TODO(), bson.M{"fed_id": fed_id})
	if ff.Err() != nil {
		return false, ""
	}
	var document bson.M
	ff.Decode(&document)
	fbanned = document["fbans"].(bson.A)
	_, IsBanned, i := deduplicate_fban(fbanned, user_id)
	if IsBanned {
		return true, fbanned[i].(bson.M)["reason"].(string)
	} else {
		return false, ""
	}
}

func Search_fed_by_id(fed_id string) bson.M {
	fed := feds.FindOne(context.TODO(), bson.M{"fed_id": fed_id})
	if fed.Err() != nil {
		return nil
	}
	var fed_info bson.M
	fed.Decode(&fed_info)
	return fed_info
}
