<template>
  <div
    :class="['modal-wrapper', { 'active': open }]">
    <span
      class="modal-overlay"
      @click="toggle"></span>

    <div class="modal">
      <div
        :class="['modal__head', { 'u-d-block u-text-center': isShowCenteredHeader }]">
        <h3 class="headline--sm">{{ header }}</h3>
        <button
          class="modal__close btn btn--icon"
          @click="toggle">
          X
        </button>
      </div>

      <div class="modal__body">
        <slot name="modal-body"></slot>
      </div>

      <div
        v-if="isShowFooterDialog"
        class="modal__footer">
        <button
          class="btn btn--lg"
          @click="toggle">{{ cancelButtonDialogText }}</button>
        <button
          @click="$emit('confirm-dialog')"
          class="btn btn--lg btn--red-500 u-ml-8">{{ confirmButtonDialogText }}</button>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  name: 'Modal',

  props: {
    open: {
      type: Boolean,
      default: false
    },

    header: {
      type: String,
      require: true
    },

    cancelButtonDialogText: {
      type: String,
      default: 'Cancel'
      // default: function () {
      //   return this.$t('modal.cancel')
      // }
    },

    confirmButtonDialogText: {
      type: String,
      default: 'Confirm'
    },

    isHiddenButtonConfirm: {
      type: Boolean,
      default: true
    },

    isShowFooterDialog: {
      type: Boolean,
      default: true
    },

    isShowCenteredHeader: {
      type: Boolean,
      default: false
    }
  },

  methods: {
    toggle() {
      console.log('object')
      this.$emit('update:open', !this.open)
    }
  }
}
</script>

<style lang="scss" scoped>
.modal-wrapper,
.modal-overlay {
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
}

.modal-wrapper {
  position: fixed;
  background-color: rgba(51, 51, 51, 0.4);
  z-index: 200;
  display: none;

  &.active {
    display: flex;
    justify-content: center;
    align-items: center;
  }
}

.modal-overlay {
  position: absolute;
}

.modal {
  border-radius: 8px;
  background-color: white;
  margin: auto 16px;
  position: relative;
  width: 100%;

  @media (min-width: 768px) {
    width: 600px;
  }

  &--lg {
    width: 800px;
  }

  &--md {
    width: 640px;
  }

  &--sm {
    width: 480px;
  }

  &__head {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 16px 24px;
    border-bottom: 1px solid #d8d8d8;
  }

  &__padding {
    height: 48px;
  }

  &__body {
    padding: 16px 24px;
    max-height: 80vh;

    &.is-scroll {
      overflow-y: auto;
    }

    &--sm {
      max-width: 348px;
      margin: auto;
    }
  }

  &__body-inner {
    padding-top: 16px;
    padding-bottom: 16px;
  }

  &__body-inner + &__body-inner {
    border-top: 1px solid #d8d8d8;
  }

  &__footer {
    font-size: 0;
    text-align: right;
    padding: 16px 24px;
    border-top: 1px solid #d8d8d8;
  }

  &__close {
    position: absolute;
    top: 8px;
    right: 8px;
    font-size: 0;
    padding: 4px;
    border-radius: 8px;
    cursor: pointer;
    color: #555;
  }

  &__action + &__action {
    margin-left: 16px;

    // @include respond-only(xs) {
    //   margin-left: 2px;
    // }
  }
}

</style>
