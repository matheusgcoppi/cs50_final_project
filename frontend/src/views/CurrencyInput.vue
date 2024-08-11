<template>
  <input
      ref="inputRef"
      type="text"
      class="currency-input"
  />
</template>

<script>
import { useCurrencyInput } from 'vue-currency-input'

export default {
  name: 'CurrencyInput',
  props: {
    modelValue: Number,
    options: Object
  },
  setup(props) {
    const { inputRef } = useCurrencyInput(props.options)

    return { inputRef }
  }
}

</script>

<style scoped>
.currency-input {
  border: 1px solid #B1B1B1; /* Add a border all the time */
  border-radius: 3px;
  text-align: left;
  padding-right: 3px;
  font-size: inherit;
  font-family: inherit;
}
</style>