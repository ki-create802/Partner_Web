<!-- src/components/Carousel.vue -->
<template>
    <div class="carousel">
      <div class="slides" :style="{ transform: `translateX(-${currentIndex * 100}%)` }">
        <div v-for="(image, index) in images" :key="index" class="slide">
          <img :src="image" alt="Slide" />
        </div>
      </div>
      <button class="prev" @click="prevSlide">&lt;</button>
      <button class="next" @click="nextSlide">&gt;</button>
    </div>
  </template>
  
  <script>
  export default {
    name: 'CarouselImage',
    props: {
      images: {
        type: Array,
        required: true
      },
      interval: {
        type: Number,
        default: 3000 // 默认间隔时间为 3 秒
      }
    },
    data() {
      return {
        currentIndex: 0,
        timer: null
      };
    },
    mounted() {
      this.startAutoPlay();
    },
    beforeUnmount() {
      this.stopAutoPlay();
    },
    methods: {
      prevSlide() {
        this.currentIndex = (this.currentIndex - 1 + this.images.length) % this.images.length;
      },
      nextSlide() {
        this.currentIndex = (this.currentIndex + 1) % this.images.length;
      },
      startAutoPlay() {
        this.timer = setInterval(() => {
          this.nextSlide();
        }, this.interval);
      },
      stopAutoPlay() {
        clearInterval(this.timer);
      }
    }
  };
  </script>
  
  <style scoped>
  .carousel {
    position: relative;
    width: 100%;
    overflow: hidden;
  }
  
  .slides {
    display: flex;
    transition: transform 0.5s ease-in-out;
  }
  
  .slide {
    flex: 0 0 100%;
    display: flex;
    justify-content: center;
    align-items: center;
  }
  
  .slide img {
    max-width: 100%;
    max-height: 100%;
  }
  
  button{
    position: absolute;
    top: 50%;
    transform: translateY(-50%);
    background-color: transparent;
    color:transparent;
    border: none;
    padding: 10px;
    cursor: pointer;
  }
  button:hover{
    color: rgb(0, 0, 0);
  }
  .prev {
    left: 0;
  }
  
  .next {
    right: 0;
  }
  </style>