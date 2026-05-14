# CSS Stylesheet Documentation

This README explains every selector and CSS property used in the stylesheet.

---

# Global Variables

## `:root`

Global CSS variables available throughout the application.

| Variable | Description |
|---|---|
| `--bg0` | Primary dark background color |
| `--bg1` | Secondary dark background color |
| `--card` | Transparent white overlay for cards |
| `--card2` | Slightly brighter card overlay |
| `--text` | Main text color |
| `--muted` | Secondary muted text color |
| `--faint` | Low-emphasis text color |
| `--line` | Border/divider color |
| `--shadow` | Reusable shadow effect |
| `--mono` | Monospace font stack |
| `--sans` | Sans-serif font stack |

---

# Universal Selector

## `*`

Applies styles to all elements.

| Property | Description |
|---|---|
| `box-sizing: border-box` | Includes padding and borders inside width/height calculations of content area. So if CA is 100px, if it has padding and border, it stay 100px |

---

# HTML & Body

## `html, body`

| Property | Description |
|---|---|
| `height: 100%` | Makes the page occupy full viewport height |

---

# Body Styling

## `body`

| Property | Description |
|---|---|
| `margin: 0` | Removes browser default margin |
| `color` | Default text color |
| `font-family` | Default font family |
| `background` | Multi-layer gradient background |

---

# Background Overlay

## `.bg`

Decorative glowing background layer.

| Property | Description |
|---|---|
| `position: fixed` | Fixed relative to viewport |
| `inset: 0` | Covers entire screen |
| `background` | Radial glow effects |
| `pointer-events: none` | Prevents mouse interaction |
| `mix-blend-mode: screen` | Blends glow with background |

---

# Main Wrapper

## `.wrap`

Main page container.

| Property | Description |
|---|---|
| `width` | Responsive container width, so it fit mobile screen also |
| `margin` | Centers container |
| `display: flex` | Uses Flexbox |
| `flex-direction: column` | Stacks sections vertically |
| `gap` | Space between children |

---

# Header

## `.header`

Top header container.

| Property | Description |
|---|---|
| `display: flex` | Flex layout |
| `align-items` | Vertical alignment |
| `justify-content` | Horizontal spacing |
| `gap` | Space between elements |

---

# Brand Section

## `.brand`

Brand/logo grouping.

| Property | Description |
|---|---|
| `display: flex` | Flex layout |
| `gap` | Space between logo and title |
| `align-items` | Vertical alignment |

---

# Logo

## `.logo`

Logo styling.

| Property | Description |
|---|---|
| `width` | Logo width |
| `height` | Logo height |
| `border-radius` | Rounded corners |
| `display: grid` | Grid layout |
| `place-items: center` | Centers content |
| `background` | Gradient background |
| `color` | Text color |
| `font-weight` | Bold text |
| `letter-spacing` | Character spacing |
| `box-shadow` | Glow/shadow effect |

---

# Titles

## `.titles .title`

Main title styling.

| Property | Description |
|---|---|
| `font-weight` | Extra bold text |
| `font-size` | Title size |
| `letter-spacing` | Character spacing |

---

## `.titles .subtitle`

Subtitle styling.

| Property | Description |
|---|---|
| `color` | Muted text color |
| `font-size` | Subtitle size |
| `margin-top` | Space above subtitle |

---

# Status Box

## `.status`

Status display container.

| Property | Description |
|---|---|
| `text-align: right` | Aligns text to right |
| `padding` | Inner spacing |
| `border-radius` | Rounded corners |
| `background` | Card background |
| `border` | Subtle border |
| `backdrop-filter: blur(10px)` | Glassmorphism blur effect |

---

## `.statusLabel`

| Property | Description |
|---|---|
| `color` | Faint text color |
| `font-size` | Small font size |

---

## `.statusCode`

| Property | Description |
|---|---|
| `display: flex` | Flex layout |
| `gap` | Space between items |
| `align-items` | Vertical alignment |
| `justify-content` | Right alignment |
| `margin-top` | Space above |

---

# Pills

## `.pill`

Rounded badge/pill.

| Property | Description |
|---|---|
| `font-family` | Monospace font |
| `font-size` | Small text |
| `padding` | Inner spacing |
| `border-radius: 999px` | Fully rounded pill shape |
| `background` | Semi-transparent background |
| `border` | Subtle border |

---

# Status Variants

## `.code-200 .pill`

OK style.

| Property | Description |
|---|---|
| `background` | Green background |
| `border-color` | Green border |

---

## `.code-202 .pill`

Accepted style.

| Property | Description |
|---|---|
| `background` | Blue background |
| `border-color` | Blue border |

---

## `.code-400 .pill`

Bad Request style.

| Property | Description |
|---|---|
| `background` | Red background |
| `border-color` | Red border |

---

# Panels

## `.panel`

Card/panel container.

| Property | Description |
|---|---|
| `background` | Transparent gradient |
| `border` | Subtle border |
| `border-radius` | Rounded corners |
| `box-shadow` | Shadow effect |
| `overflow: hidden` | Hides overflow |

---

# Forms

## `.form`

| Property | Description |
|---|---|
| `padding` | Inner spacing |

---

# Modes

## `.modes`

Mode selector container.

| Property | Description |
|---|---|
| `display: flex` | Flex layout |
| `gap` | Space between items |
| `margin-bottom` | Space below |

---

# Radio Buttons

## `.radio`

Custom radio container.

| Property | Description |
|---|---|
| `display: flex` | Flex layout |
| `gap` | Space between elements |
| `align-items` | Vertical alignment |
| `padding` | Inner spacing |
| `border-radius` | Rounded corners |
| `border` | Subtle border |
| `background` | Transparent background |

---

## `.radio input`

| Property | Description |
|---|---|
| `accent-color` | Changes selected radio color |

---

## `.radio span`

Custom Decode/Encode text

| Property | Description |
|---|---|
| `font-size` | Text size |
| `color` | Muted text color |

---

# Labels

## `.label`

| Property | Description |
|---|---|
| `display: block` | Makes label full width |
| `margin` | Space around label |
| `color` | Muted text color |
| `font-size` | Small text size |

---

# Inputs

## `.input`

| Property | Description |
|---|---|
| `width: 100%` | Full width input |
| `padding` | Inner spacing |
| `border-radius` | Rounded corners |
| `border` | Subtle border |
| `background` | Dark transparent background |
| `color` | Text color |
| `font-family` | Monospace font |
| `line-height` | Line spacing |
| `outline: none` | Removes browser outline |

---

## `.input:focus`

Focused input state.

| Property | Description |
|---|---|
| `border-color` | Blue focus border |
| `box-shadow` | Glow effect |

---

# Actions

## `.actions`

Button/action row.

| Property | Description |
|---|---|
| `margin-top` | Space above |
| `display: flex` | Flex layout |
| `align-items` | Vertical alignment |
| `justify-content` | Space between items |
| `gap` | Space between items |
| `flex-wrap: wrap` | Allows wrapping |

---

# Buttons

## `.btn`

Primary button styling.

| Property | Description |
|---|---|
| `appearance: none` | Removes native styling |
| `border: 0` | Removes border |
| `cursor: pointer` | Pointer cursor |
| `padding` | Inner spacing |
| `border-radius` | Rounded corners |
| `color` | Text color |
| `font-weight` | Bold text |
| `letter-spacing` | Character spacing |
| `background` | Gradient background |
| `box-shadow` | Glow/shadow effect |

---

## `.btn:active`

Pressed state.

| Property | Description |
|---|---|
| `transform: translateY(1px)` | Moves button slightly down |

---

# Hints

## `.hint`

| Property | Description |
|---|---|
| `color` | Faint text color |
| `font-size` | Small text size |

---

## `.hint code`

Inline code styling.

| Property | Description |
|---|---|
| `font-family` | Monospace font |
| `color` | Bright text |
| `background` | Transparent background |
| `border` | Subtle border |
| `padding` | Inner spacing |
| `border-radius` | Rounded corners |

---

# Output Panel

## `.outputPanel`

| Property | Description |
|---|---|
| `padding` | Inner spacing |

---

## `.outputHeader`

| Property | Description |
|---|---|
| `display: flex` | Flex layout |
| `align-items` | Vertical alignment |
| `justify-content` | Space between items |
| `gap` | Space between items |
| `margin-bottom` | Space below |

---

## `.outputTitle`

| Property | Description |
|---|---|
| `font-weight` | Bold title |
| `letter-spacing` | Character spacing |

---

# Status Messages

## `.error`

Error badge styling.

| Property | Description |
|---|---|
| `color` | Red text |
| `font-size` | Small text |
| `background` | Transparent red background |
| `border` | Red border |
| `padding` | Inner spacing |
| `border-radius` | Rounded badge |

---

## `.ok`

Success badge styling.

| Property | Description |
|---|---|
| `color` | Green text |
| `font-size` | Small text |
| `background` | Transparent green background |
| `border` | Green border |
| `padding` | Inner spacing |
| `border-radius` | Rounded badge |

---

# Art / Code Display

## `.art`

Code or ASCII art display area.

| Property | Description |
|---|---|
| `margin` | Removes outer spacing |
| `padding` | Inner spacing |
| `border-radius` | Rounded corners |
| `border` | Subtle border |
| `background` | Dark background |
| `font-family` | Monospace font |
| `font-size` | Text size |
| `line-height` | Line spacing |
| `min-height` | Minimum height |
| `overflow: auto` | Enables scrolling |
| `white-space: pre` | Preserves formatting |

---

# Links

## `.link`

| Property | Description |
|---|---|
| `color` | Blue link color |
| `text-decoration: none` | Removes underline |
| `font-weight` | Medium bold text |

---

## `.link:hover`

Hover state.

| Property | Description |
|---|---|
| `text-decoration: underline` | Adds underline on hover |

---

# Footer

## `.footer`

Footer text styling.

| Property | Description |
|---|---|
| `color` | Faint text color |
| `font-size` | Small text size |
| `text-align: center` | Centers text |
| `margin-top` | Space above |

---
