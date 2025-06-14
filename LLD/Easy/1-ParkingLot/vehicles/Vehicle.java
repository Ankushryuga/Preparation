public class Vehicle{
    protected String license;
    protected VehicleTypes vehicleType;

    public Vehicle(String license, VehicleTypes vehicleType){
        this.license=license;
        this.vehicleType=vehicleType;
    }

    public VehicleTypes getVehicleTypes(){
        return vehicleType;
    }
    //getter and setter
}