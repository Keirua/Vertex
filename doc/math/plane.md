# Plane intersection equation

http://homeomath2.imingo.net/equplan.htm

One of the basic primitives is the infinite plane P, which is defined by 2 components : 

 - a point on the plane A (xA, yA, zA)
 - a normal n (xN, yN, zN)

M(x,y,z) € P(A, n)

<=> AM.n = 0 <=> dot(AM, n) = 0
<=> (x-xA, y-yA, z-zA).(nX, nY, nZ) = 0
<=> nX*(x-xA)+nY*(y-yA)+nZ*(z-zA) = 0
<=> nX*x + nY*y + nZ*z -(nX*xA+nY*yA+nZ*Z) = 0

let D = -(nX*xA+nY*yA+nZ*Z) = -dot(n, OA)

If we replace the values for x, y, y with the one of our ray (origin O, direction Dir), we can solve the equation in t. After simplification, we get the following value for t :

t = - ( (nX*X + Y*Y + nZ*Z + D) / (nX*DIR.x + nY*DIR.y + nZ*DIR.z) )
with (X Y Z) = (O.x-xA, O.y-yA, O.z-zA)

Note that if (dot(n, DIR) == 0), the ray is parallel to the plane : there is no solution (and the denominator is nul)

As for the normal at the intersection point, that's the normal of the plane