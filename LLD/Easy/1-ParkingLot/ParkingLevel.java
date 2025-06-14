
import java.util.List;
import java.util.Optional;

public class ParkingLevel{
    int parkingLevel;
    int capacity;
    List<ParkingSlot> parkingSlots;
    public List<ParkingSlot> getParkingSlots(){
        return parkingSlots;
    }

    public synchronized Optional<ParkingSlot> getAvailableSpots(VehicleTypes types){
return parkingSlots.stream()
                .filter(spot -> spot.isAvailable() && spot.getVehicleTypes() == types)
                .findFirst();
    }

    public int getFloorNumber(){
        return parkingLevel;
    }

}