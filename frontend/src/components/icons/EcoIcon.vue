<template>
  <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1200 300">
    <!-- Emblem: circular leaf -->
    <g transform="translate(100,150)">
      <circle cx="0" cy="0" r="80" class="line" />
      <path
        class="leaf"
        d="M0,-60 C36,-60 60,-24 60,0 C60,48 0,120 0,120 C0,120 -60,48 -60,0 C-60,-24 -36,-60 0,-60 Z"
      />
    </g>

    <!-- Brand Name -->
    <text x="250" y="150" class="text" dominant-baseline="middle">eco quest calculator</text>

    <!-- Decorative baseline -->
    <path d="M250,200 L1100,200" class="line" />
  </svg>
</template>

<style scoped>
.text {
  font-family: 'Poppins', sans-serif;
  font-weight: 600;
  font-size: 100px;
  fill: #2e7d32;
  letter-spacing: 3px;
}

.line {
  stroke: #76c043;
  stroke-width: 3;
  fill: none;
}

.icon {
  stroke: #2e7d32;
  stroke-width: 3;
  fill: none;
  stroke-linecap: round;
  stroke-linejoin: round;
}

.leaf {
  fill: #a5d6a7;
}
</style>
