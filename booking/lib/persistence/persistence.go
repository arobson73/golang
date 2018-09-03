package persistence

type DatabaseHandler interface {
	AddUser(User) ([]byte, error)
	AddAdminUser(AdminUser) ([]byte, error)
	AddEvent(Event) ([]byte, error)
	AddEventForAdminUser([]byte, Event) error
	AddBookingForUser([]byte, Booking) error
	AddLocation(Location) (Location, error)
	FindUser(string, string) (User, error)
	FindUserEmailPass(string, string) (User, error)
	FindAdminUserEmailPass(string, string) (AdminUser, error)
	FindBookingsForUser([]byte) ([]Booking, error)
	FindEvent([]byte) (Event, error)
	FindAdminUser([]byte) (AdminUser, error)
	FindEventByName(string) (Event, error)
	FindAllAvailableEvents() ([]Event, error)
	FindLocation(string) (Location, error)
	FindAllLocations() ([]Location, error)
}
