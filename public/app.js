const publicVapidKey = "YOUR_PUBLIC_VAPID_KEY";

if ("serviceWorker" in navigator) {
    registerServiceWorker().catch(console.error);
}

async function registerServiceWorker() {
    const register = await navigator.serviceWorker.register("/worker.js");
    console.log("Service Worker Registered.");

    // Subscribe for Push Notifications
    const subscription = await register.pushManager.subscribe({
        userVisibleOnly: true,
        applicationServerKey: urlBase64ToUint8Array(publicVapidKey),
    });

    console.log("Push Notification Subscription:", subscription);

    // Send subscription to the server
    await fetch("/subscribe", {
        method: "POST",
        body: JSON.stringify(subscription),
        headers: {
            "Content-Type": "application/json",
        },
    });
}

function urlBase64ToUint8Array(base64String) {
    const padding = "=".repeat((4 - (base64String.length % 4)) % 4);
    const base64 = (base64String + padding).replace(/-/g, "+").replace(/_/g, "/");
    const rawData = window.atob(base64);
    return Uint8Array.from([...rawData].map((char) => char.charCodeAt(0)));
}
