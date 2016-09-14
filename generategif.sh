#!/bin/sh

convert -delay 1000 -loop 0 out/1-out-antialiasing.png \
out/2-out-lighting-as3.png \
out/3-out-2lights-softershadows-as3.png \
out/4-out-2lights-better-settings.png \
out/5-out-reflection-2lights-depth2.png \
out/6-out-texture-checkboard-depth1-as5.png \
out/7-out-marble-turbulence-depth1-as3.png \
out/8-out-blinnphong-specular-hightlights-depth1-as3.png \
out/9-out-exponential-exposure--1.66-depth1-as3.png \
out/10-out-soft-shadows-16rays-0.2strength-as3-d2-exposure1.66.png \
out/animation.gif


