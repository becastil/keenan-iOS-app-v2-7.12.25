# Sydney Health Android App

Android application built with Kotlin and RIBs architecture, replicating the functionality of Anthem Sydney Health app.

## Architecture

This app uses Uber's RIBs (Router, Interactor, Builder) architecture pattern for Android with:
- Clear separation of concerns
- Testability
- Deep linking support
- Business logic isolation
- View-independent business logic

## Tech Stack

- **Language**: Kotlin
- **Architecture**: RIBs
- **UI**: Jetpack Compose + XML Views
- **Reactive**: RxJava 3 + Coroutines
- **Networking**: Retrofit + OkHttp + gRPC
- **DI**: Dagger 2
- **Image Loading**: Coil

## Requirements

- Android Studio Hedgehog (2023.1.1) or newer
- Android SDK 34
- Minimum SDK 24 (Android 7.0)
- JDK 17

## Setup

1. Open the project in Android Studio:
```bash
cd android
studio .
```

2. Sync the project with Gradle files

3. Run the app on an emulator or device

## Project Structure

```
app/src/main/
├── java/com/uber/sydneyhealth/
│   ├── ribs/                # RIBs components
│   │   ├── root/           # Root RIB
│   │   ├── loggedout/      # Login flow
│   │   ├── loggedin/       # Main app flow
│   │   ├── dashboard/      # Dashboard RIB
│   │   ├── benefits/       # Benefits RIB
│   │   ├── claims/         # Claims RIB
│   │   ├── providers/      # Provider search RIB
│   │   ├── membercard/     # Member ID RIB
│   │   └── messages/       # Messaging RIB
│   ├── services/           # Business logic services
│   ├── models/             # Data models
│   └── utils/              # Utilities and extensions
├── res/                    # Resources
│   ├── layout/            # XML layouts
│   ├── values/            # Strings, colors, styles
│   └── drawable/          # Icons and graphics
└── AndroidManifest.xml
```

## RIBs Architecture

Each RIB consists of:
- **Builder**: Creates and configures the RIB
- **Router**: Manages child RIBs and navigation
- **Interactor**: Contains business logic
- **View**: UI components (optional)
- **Presenter**: Formats data for display (optional)

## Key Features

- Biometric authentication (Fingerprint/Face)
- Offline support with Room database
- Push notifications via Firebase
- Deep linking support
- Real-time messaging
- Secure credential storage
- Material 3 design system

## Build Variants

- **debug**: Development build with debugging enabled
- **release**: Production build with ProGuard/R8

## Testing

Run unit tests:
```bash
./gradlew test
```

Run instrumented tests:
```bash
./gradlew connectedAndroidTest
```

## Code Style

We use Kotlin coding conventions with ktlint for enforcement.

## Security

- Encrypted SharedPreferences for sensitive data
- Certificate pinning for API calls
- ProGuard rules for code obfuscation

## Authentication

Demo credentials:
- Member ID: M123456
- Password: demo

## Performance

- Lazy loading of RIBs
- Image caching with Coil
- Efficient list rendering with RecyclerView
- Background task optimization

## Accessibility

- Content descriptions for all interactive elements
- Support for TalkBack
- High contrast mode support
- Dynamic font sizing