import java.util.*;
import java.util.concurrent.ConcurrentHashMap;

public class ParkingLot{
    private static ParkingLot instance;
    private List<ParkingLevel> floors;

    Map<String, Ticket> activeTickets=new ConcurrentHashMap<>();        //its for inmemory.

    public ParkingLot(){
        floors=new ArrayList<>();
    }

    public static synchronized ParkingLot getInstance(){
        if(instance==null){
            instance=new ParkingLot();
        }
        return instance;
    }

    public void addFloor(ParkingLevel level){
        floors.add(level);
    }

    public synchronized Ticket parkVehicle(Vehicle vehicle) throws Exception{
        for(ParkingLevel level:floors){
//create logic.
        }
        return null;
    }

}