import React, { useEffect, useState } from 'react';

const RentHistory = () => {
    const [rentals, setRentals] = useState([]);
    const userID = 1; // Replace with actual user ID

    useEffect(() => {
        const fetchRentals = async () => {
            const response = await fetch(`http://localhost:8082/rent/history?user_id=${userID}`);
            const data = await response.json();
            setRentals(data);
        };
        fetchRentals();
    }, [userID]);

    return (
        <div>
            <h1>Rental History</h1>
            <ul>
                {rentals.map((rental) => (
                    <li key={rental.id}>{rental.car_id} from {rental.start_date} to {rental.end_date}</li>
                ))}
            </ul>
        </div>
    );
};

export default RentHistory;
