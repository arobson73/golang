
export interface Event {
    ID: string;
    Name: string;
    StartDate: number;
    EndDate: number;
    Location: {
        ID: string;
        Name: string;
    }
}

export interface Booking {
    Seats: number;
    Date:number;
    EventID:string;
}
