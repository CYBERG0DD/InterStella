# Media Assets Guide

## Required Media Files

To complete the visual experience, add the following media files to your project:

### Video Files (16:9 cinematic ratio)
Place in `/public/media/` directory:

- `phase-01.mp4` - Ignition phase video
- `phase-02.mp4` - Foundation phase video
- `phase-03.mp4` - Core System phase video
- `phase-04.mp4` - Convergence phase video
- `phase-05.mp4` - Expansion phase video
- `phase-06.mp4` - Intelligence phase video
- `phase-07.mp4` - Dominion phase video

### Image Files (4:3 ratio recommended)
Place in `/public/media/` directory:

- `phase-01.jpg` - Ignition phase image
- `phase-02.jpg` - Foundation phase image
- `phase-03.jpg` - Core System phase image
- `phase-04.jpg` - Convergence phase image
- `phase-05.jpg` - Expansion phase image
- `phase-06.jpg` - Intelligence phase image
- `phase-07.jpg` - Dominion phase image

## How to Add Media

1. Create a `/public/media/` directory in your project
2. Add your video and image files with the exact names above
3. Uncomment the `<video>` and `<img>` tags in `/src/app/components/PhaseSection.tsx`

## Recommended Specifications

### Videos
- Format: MP4 (H.264)
- Resolution: 1920x1080 (Full HD)
- Duration: 10-30 seconds
- Loop-friendly (seamless beginning/end)
- Size: Keep under 5MB per video for optimal loading

### Images
- Format: JPG or WebP
- Resolution: 1600x1200 or higher
- Optimized for web (compressed)
- Size: Keep under 500KB per image

## Alternative: Using Placeholder Services

If you need temporary placeholders, you can update the URLs in `/src/app/data/phases.ts`:

```typescript
videoUrl: "https://your-video-cdn.com/phase-01.mp4"
imageUrl: "https://images.unsplash.com/your-image-id"
```

## Current State

The application currently displays styled placeholder boxes for each media element. Once you add the actual files and uncomment the media tags, they will automatically appear with smooth animations and hover effects.
