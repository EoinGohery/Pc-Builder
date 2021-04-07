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
    private long id;
    @NotBlank(message = "Name cannot be empty or Null")
    private String gpuName;
    private String Manufacturer;
    private String Memory;
    private String clock;
    private Integer TDP;
    private Integer price;
}
