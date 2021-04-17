package com.example.BriarfieldPC.demo.entity;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import javax.validation.constraints.NotBlank;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;

import static javax.persistence.GenerationType.IDENTITY;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Builder
@Entity
public class Gpu {
    @Id
    @GeneratedValue(strategy = IDENTITY)
    public long Id;
    public String Manufacturer;
    @NotBlank(message = "Name cannot be empty or Null")
    public String GpuName;
    public String Memory;
    public String Clock;
    public Integer Tdp;
    public Integer Price;
}
