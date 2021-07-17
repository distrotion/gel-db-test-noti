# minefit_noti_demo_07-21

#### input test

mutation{
  createNoti(input:{
    userid: "",
  	name: "",
    related: "",
  	type: "",
  	doctor: "",
  	appointment: "",
  })
}

mutation{
  callNoti(input:{
    userid: "1-60ef67f01691c975f04e7550"
  }){
    notilist{
    name
    type
    related
    doctor
    appointment
  }
  }
}



