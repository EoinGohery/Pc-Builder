package com.example.BriarfieldPC.demo.entity;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;
import javax.validation.constraints.NotBlank;

import static javax.persistence.GenerationType.IDENTITY;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Builder
@Entity
public class PSU {
    @Id
    @GeneratedValue(strategy = IDENTITY)
    private long id;
    private String name;
    private Integer price;
    private String Manufacturer;
    private Integer capacity;
    private Integer rating;
    private Integer TDP;
}
