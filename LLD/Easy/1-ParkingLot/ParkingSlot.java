public class ParkingSlot{
    private int slotId;
    private boolean isAvailable;
    private Vehicle vehicle;
    private VehicleTypes vehicleTypes;
    public ParkingSlot(int slotId, boolean isAvailable){
        this.slotId=slotId;
        this.isAvailable=false; 
    }
    public synchronized boolean isAvailable(){
        return !isAvailable;
    }
    public synchronized boolean park(Vehicle vehicle){
        if(isAvailable){
            return false;
        }
        this.vehicle=vehicle;
        isAvailable=true;
        return true;
    }
    public synchronized void unpark(){
        vehicle=null;
        isAvailable=false;
    }
    public VehicleTypes getVehicleTypes(){
        return vehicleTypes;
    }
    public Vehicle getVehicle(){
        return vehicle;
    }
}