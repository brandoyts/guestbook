"use client"

import React, { ChangeEvent, FormEvent, useState } from "react"

import Api from "@/services/api"
import { Button } from "@nextui-org/button"
import { Input } from "@nextui-org/input"

function Form() {
  const [inputValue, setInputValue] = useState("")

  const handleInputChange = (e: ChangeEvent<HTMLInputElement>) => {
    setInputValue(e.target.value)
  }

  const handleSubmit = async (e: FormEvent) => {
    try {
      e.preventDefault()
      Api.setItemToRedis(inputValue)
    } catch (error) {
      throw error
    }
  }

  return (
    <form action="#" onSubmit={handleSubmit} className="flex gap-2">
      <Input
        placeholder="type something..."
        value={inputValue}
        onChange={handleInputChange}
        className="color-red-500"
      />
      <Button color="primary" type="submit">
        Submit
      </Button>
    </form>
  )
}

export default Form
