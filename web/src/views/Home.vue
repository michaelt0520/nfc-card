<template>
  <div>
    <Header />
    <sign-up-form />
    <price-table />
    <contact-form />
  </div>
</template>

<script>
import Header from "../components/Header";
import SignUpForm from "../components/SignUpForm";
import PriceTable from "../components/PriceTable";
import ContactForm from "../components/ContactForm";

export default {
  name: "Home",

  components: {
    Header,
    SignUpForm,
    PriceTable,
    ContactForm,
  },
};
</script>
