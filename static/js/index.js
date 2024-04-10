const Services = [
    { name: 'Game Service', id: 'game_backend' },
    { name: 'Game Oppenent',  id: 'game_ai' },
    { name: 'User service', id: 'game_auth' },
    { name: 'Front End', id: 'game_client' }
];

async function toggleService(id, value) {
    try {
        const response = await fetch(`http://localhost:1234/${value ? 'start' : 'stop'}/${id}`, {
            method: 'POST'
        });
        const result = await response.text();
        console.log("🚀 ~ toggleService ~ response:", result)

        if (!response.ok) {
            throw new Error('Failed to toggle service');
        }
        const button = document.getElementById(`${id}-toggle`);
        button.textContent = value ? 'ON' : 'OFF';
    } catch (error) {
        console.error('Error toggling service:', error);
    }
}

async function handleCheckHealth(id) {
    try {
        const response = await fetch(`http://localhost:1234/health/${id}`);
        if (!response.ok) {
            throw new Error('Failed to fetch health status');
        }
        const result = await response.json();
        console.log("🚀 ~ handleCheckHealth ~ result:", result)
        document.getElementById(`${id}-health-status`).textContent = result.healthy ? 'Healthy' : 'Unhealthy' ;
    } catch (error) {
        console.error('Error checking health:', error);
    }
}
function init() {
    const container = document.getElementById('container');

    Services.forEach(service => {
        const div = document.createElement('div');
        div.classList.add('service');

        const serviceName = document.createElement('span');
        serviceName.textContent = service.name;

        const toggleButton = document.createElement('button');
        toggleButton.textContent = 'OFF';
        toggleButton.id = `${service.id}-toggle`;
        toggleButton.onclick = () => toggleService(service.id, toggleButton.textContent === 'OFF');

        const checkButton = document.createElement('button');
        checkButton.textContent = 'Check Health';
        checkButton.onclick = () => {
            handleCheckHealth(service.id)
        };

        const healthDescription = document.createElement('span');
        healthDescription.classList.add('health-status');
        healthDescription.id = `${service.id}-health-status`;
        healthDescription.textContent = 'Health status';

        div.appendChild(serviceName);
        div.appendChild(toggleButton);
        div.appendChild(checkButton);
        div.appendChild(healthDescription);
        container.appendChild(div);
    });
}

window.onload = init;