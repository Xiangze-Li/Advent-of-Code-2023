# Math for Puzzle 24 Part 2

For hailstone $i$, assume the initial point is $\vec{p}_i$, initial velocity is $\vec{v}_i$.

Assume that out stone have initial point $\vec{p}_0$ and velocity $\vec{v}_0$.

Assume out stone collides with hailstone $i$ at time $t_i$.

That makes

$$
\vec{p}_i + \vec{v}_i t_i = \vec{p}_0 + \vec{v}_0 t_i
$$

Then

$$
(\vec{p}_0 - \vec{p}_i) + (\vec{v}_0 - \vec{v}_i) t_i = \vec0
$$

For $t_i$ is a scalar, vector $(\vec{p}_0 - \vec{p}_i)$ is parallel to $(\vec{v}_0 - \vec{v}_i)$. Their cross product is zero vector

$$
\begin{aligned}
(\vec{p}_0 - \vec{p}_i) \times (\vec{v}_0 - \vec{v}_i) &= \vec0 \\
\vec{p}_0 \times \vec{v}_0 -
\vec{p}_0 \times \vec{v}_i -
\vec{p}_i \times \vec{v}_0 +
\vec{p}_i \times \vec{v}_i &= 0
\end{aligned}
$$

The $\vec{p}_0 \times \vec{v}_0$ part makes this a bi-linear equation.

However, this part is the same for each $i$. Taking a pair of diffrent $i$ and substract the two equation cancels this out.

$$
\begin{array}{rlc}
\vec{p}_0 \times \vec{v}_0 -
\vec{p}_0 \times \vec{v}_i -
\vec{p}_i \times \vec{v}_0 +
\vec{p}_i \times \vec{v}_i &= 0 &(a) \\
\vec{p}_0 \times \vec{v}_0 -
\vec{p}_0 \times \vec{v}_j -
\vec{p}_j \times \vec{v}_0 +
\vec{p}_j \times \vec{v}_j &= 0 &(b) \\

- \vec{p}_0 \times (\vec{v}_i - \vec{v}_j) -
(\vec{p}_i - \vec{p}_j) \times \vec{v}_0 +
\vec{p}_i \times \vec{v}_i - \vec{p}_j \times \vec{v}_j &= 0 &(a)-(b)
\end{array}
$$

We have two unknown vector $\vec{p}_0$ and $\vec{v}_0$, then we need two pairs of $i$ to make two linear equations. Or, speaking of unknown component, two pairs to make six scalar linear equations.

(The six equations omitted)
