---
title: WRBC App Now Live
date: 2025-01-04
draft: true
---

Our club website is now available as a mobile app! Using it as an app has some
benefits:

1. If you long press on the app, it provides quick access to parts of the site members
access the most: digital membership card (requires occasional google auth); gauntlet schedule; and style guide
2. The site is cached for better response times and offline access
3. Future features, like push notifications

## How to install

### Android

Click the install button below and you should be guided through the installation. 
If the button doesn't work, go to the options menu and choose "Add to Home Screen"
<button id="installApp">Install</button>
<script>
  let deferredPrompt;
    window.addEventListener('beforeinstallprompt', (e) => {
        deferredPrompt = e;
    });

    const installApp = document.getElementById('installApp');
    installApp.addEventListener('click', async () => {
        if (deferredPrompt !== null) {
            deferredPrompt.prompt();
            const { outcome } = await deferredPrompt.userChoice;
            if (outcome === 'accepted') {
                deferredPrompt = null;
            }
        }
    });
</script>

### IPhone

Go to the share or bookmark action and choose "Add to home screen" 

If you have trouble, here's a tutorial https://www.lifewire.com/home-screen-icons-in-safari-for-iphone-and-amp-ipod-touch-4103654
