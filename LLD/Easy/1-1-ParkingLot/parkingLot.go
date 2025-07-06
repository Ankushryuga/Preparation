package main
import (
        "errors"
        "fmt"
        "sync"
        )

type TotalFloor int    //number of floors available for parking in a building.
type VehicleType int    //different type of vehicles.

//enum for vehicle types.
const (
  CAR VehicleType = iota
  BIKE
  TRUCK
  )

type Floor struct{
  FloorId          int
  AvailableSpace   int
  TotalSpace       int
}

type Space struct{
  SpaceId        int
  FloorId        int
  isAvailable    bool
  BookedBy       string
  VehicleType    VehicleType
}

type BookedInformation struct{
  BookingId     int
  SpaceId       int
  UserId        string
  VehicleType   VehicleType
}

type ParkingLotImplement struct{
  Floors         map[int]*Floor    //map being used for time complexity and auto indexing etc...
  Spaces         map[int]*Space
  Booking        map[int]*BookedInformation
  NextSpaceId    int
  NextBookingId  int
  mu             sync.Mutex        //shared mutex:::
}

type ParkingLot interface{
  AddMoreFloor(floor Floor) (int, int)    //it will return updated Floor count.
  BookSpace(userId string, VehicleType VehicleType) (Space, error)  //it will return Space struct information.
  GetBookedInformation(bookingId int) (BookedInformation, error)    //it will return BookedInformaiton of space.
  GetAvailableSpaces(vehicleType VehicleType)[]Space
}

func NewParkingLot() *ParkingLotImplement{
  return &ParkingLotImplement{
  Floors:          make(map[int]*Floor),
  Spaces:          make(map[int]*Space),
  Booking:         make(map[int]*BookedInformation),
  }
}

func (p *ParkingLotImplement) AddMoreFloor(floor Floor) (int, int){
        p.mu.Lock()
        defer p.mu.Unlock()
  p.Floors[floor.FloorId]  = &floor

  for i:=0;i<floor.TotalSpace;i++{
    space := &Space{
       SpaceId:    p.NextSpaceId,
       FloorId:    floor.FloorId,
       isAvailable: true,
       VehicleType:  CAR,  //default.
    }
    p.Spaces[p.NextSpaceId]=space
    p.NextSpaceId++
  }
  return len(p.Floors), len(p.Spaces)
}


func (p *ParkingLotImplement) BookSpace(userId string, vehicleType VehicleType) (Space, error){
  p.mu.Lock()
  defer p.mu.Unlock()
  for _, space := range p.Spaces{
    if space.isAvailable && space.VehicleType == vehicleType {
      
      space.isAvailable=false
      space.BookedBy=userId

      booking := &BookedInformation{
        BookingId : p.NextBookingId,
        SpaceId : space.SpaceId,
        UserId : userId,
        VehicleType : vehicleType,
      }
      p.Booking[p.NextBookingId]=booking
      p.NextBookingId++
      return *space, nil
    }
  }
  return Space{}, errors.New("No Available space for the requested vehicle type")
}


func (p *ParkingLotImplement) GetBookedInformation(bookingId int) (BookedInformation, error){
p.mu.Lock()
        p.mu.Unlock()
  if booking, exists := p.Booking[bookingId]; exists{
    return *booking, nil
  }
  return BookedInformation{}, errors.New("Booking not found")
}


func (p* ParkingLotImplement) GetAvailableSpaces(vehicleType VehicleType) []Space{
        p.mu.Lock()
        defer p.mu.Unlock()
  var available []Space
  for _, space := range p.Spaces{
    if space.isAvailable && space.VehicleType == vehicleType{
        available = append(available, *space)
    }
  }
  return available
}


func main(){
  parkingLot := NewParkingLot()
  parkingLot.AddMoreFloor(Floor{FloorId: 1, TotalSpace: 5, AvailableSpace: 5})

  //Book a space:
  space, err := parkingLot.BookSpace("user1", CAR)
  if err!=nil{
    fmt.Println("Booking error", err)
  }else{
    fmt.Println("Booking Space", space)
  }


  //Get Available Spaces::
  available := parkingLot.GetAvailableSpaces(CAR)
  fmt.Println("Available Spaces for CAR: ", available)


  //Get Booking infor::
  info, err := parkingLot.GetBookedInformation(0)
  if err==nil{
    fmt.Println("Booking info", info)
  }else{
    fmt.Println("Booking info error", err)
  }
}
