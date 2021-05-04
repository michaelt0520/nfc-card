export default {
  data() {
    return {
      debounce: null,
    }
  },

  methods: {
    handleDebounceSearch(fn, value, time) {
      clearTimeout(this.debounce)
      this.debounce = setTimeout(() => {
        fn(value)
      }, time)
    }
  }
}
