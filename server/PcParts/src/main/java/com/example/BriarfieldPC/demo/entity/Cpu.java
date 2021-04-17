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
public class Cpu {
    @Id
    @GeneratedValue(strategy = IDENTITY)
    private long id;
    private String Manufacturer;
    private String name;
    private Integer cores;
    private String clock;
    private Integer TDP;
    private String socket;
    private Integer price;
}
