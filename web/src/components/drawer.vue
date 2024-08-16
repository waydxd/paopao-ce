<template>
    <Transition name="drawer">
    <div class="drawer">
      <div class="drawer-header">
        <h2 class="drawer-title">Create Community</h2>
      </div>
      <div class="drawer-content">
        <p>This is the content of the drawer. You can add any HTML elements here.</p>
        <p>The drawer has a smooth animation and a modern, clean look.</p>
      </div>
    </div>
  </Transition>
</template>

<style scoped>
body {
  font-family: Arial, sans-serif;
  margin: 0;
  padding: 0;
  background-color: #f0f0f0;
}

@keyframes smooth-appear {
  to{
    bottom: 0;
    opacity:1;
  }
}
.drawer {
  position: absolute;
  bottom: -80%;
  width: 100%;
  height: 80%;
  background-color: white;
  box-shadow: 0 -2px 10px rgba(0, 0, 0, 0.1);
  border-top-left-radius: 20px;
  border-top-right-radius: 20px;
  padding: 20px;
  box-sizing: border-box;
  z-index: 999;
  animation: smooth-appear 0.3s ease forwards;
}

.drawer-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.drawer-title {
  font-size: 24px;
  font-weight: bold;
  color: #333;
}

.drawer-close {
  background: none;
  border: none;
  font-size: 24px;
  color: #999;
  cursor: pointer;
}

.drawer-content {
  color: #666;
  line-height: 1.6;
}
</style>
