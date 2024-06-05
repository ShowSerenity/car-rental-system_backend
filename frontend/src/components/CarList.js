import React, { useEffect, useState } from 'react';

const CarList = () => {
    const [cars, setCars] = useState([]);

    useEffect(() => {
        const fetchCars = async () => {
            const response = await fetch('http://localhost:8081/cars');
            const data = await response.json();
            setCars(data);
        };
        fetchCars();
    }, []);

    return (
        <div>
            <h1>Car List</h1>
            <ul>
                {cars.map((car) => (
                    <li key={car.id}>{car.make} {car.model}</li>
                ))}
            </ul>
        </div>
    );
};

export default CarList;
