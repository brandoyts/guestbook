import AppTitle from "@/components/app-title"
import Form from "@/components/form"

export default function Home() {
  return (
    <section className="flex flex-col items-center justify-center gap-12 py-8 md:py-10">
      <AppTitle />
      <Form />
    </section>
  )
}
