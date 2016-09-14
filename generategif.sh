#!/bin/sh

convert -quality 100 -delay 100 -loop 0 \
-fill white -pointsize 20 \
\( -annotate +10+30 "Flat shading" out/1-out-antialiasing.png \) \
\( -annotate +10+30 "Lambert lights" out/2-out-lighting-as3.png \) \
\( -annotate +10+30 "Shadows" out/3-out-2lights-softershadows-as3.png \) \
\( -annotate +10+30 "Scene update" out/4-out-2lights-better-settings.png \) \
\( -annotate +10+30 "Reflection" out/5-out-reflection-2lights-depth2.png \) \
\( -annotate +10+30 "Texture mapping" out/6-out-texture-checkboard-depth1-as5.png \) \
\( -annotate +10+30 "Random textures" out/7-out-marble-turbulence-depth1-as3.png \) \
\( -annotate +10+30 "Specular highlights" out/8-out-blinnphong-specular-hightlights-depth1-as3.png \) \
\( -annotate +10+30 "Exposure correction" out/9-out-exponential-exposure--1.66-depth1-as3.png \) \
\( -annotate +10+30 "Soft shadows" out/10-out-soft-shadows-16rays-0.2strength-as3-d2-exposure1.66.png \) \
out/features-evolution.gif